package model

import "database/sql"
import "errors"

// ErrPasswordUnmatch is error for password unmatch when logging in.
var ErrPasswordUnmatch = errors.New("password unmatch")

// UserOne returns the user for given id
func UserOne(db *sql.DB, id int64) (User, error) {
	return ScanUser(db.QueryRow(`select * from users where user_id = ?`, id))
}

// UserByEmail fetch user by email.
// Email is unique key.
func UserByEmail(db *sql.DB, email string) (User, error) {
	return ScanUser(db.QueryRow(`select * from users where email = ?`, email))
}

// UserExists check if user is exists by given email.
func UserExists(db *sql.DB, email string) (bool, error) {
	var count int64
	if err := db.QueryRow(`select count(*) as count from users where email = ?`, email).Scan(&count); err != nil {
		return false, err
	}
	return count == 1, nil
}

// Update updates user by given user.
func (u *User) Update(tx *sql.Tx) (sql.Result, error) {
	stmt, err := tx.Prepare(`
	update users
		set name = ?, email = ?, updated = CURRENT_TIMESTAMP
		where user_id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(u.Name, u.Email, u.ID)
}

// Insert inserts new user.
func (u *User) Insert(tx *sql.Tx, password string) (sql.Result, error) {
	stmt, err := tx.Prepare(`
	insert into users (name, email, salt, salted, created, updated)
	values(?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	salt := Salt(100)
	return stmt.Exec(u.Name, u.Email, salt, Stretch(password, salt))
}

// Auth makes user authentication.
func Auth(db *sql.DB, email, password string) (User, error) {
	u, err := UserByEmail(db, email)
	if err != nil {
		return User{}, err
	}
	if u.Salted != Stretch(password, u.Salt) {
		return User{}, ErrPasswordUnmatch
	}
	return u, nil
}
