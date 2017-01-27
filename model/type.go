//go:generate $GOPATH/bin/scaneo -o type_scans.go $GOFILE

package model

import "time"

// User returns model object for user.
type User struct {
	ID      int64     `json:"user_id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Salt    string    `json:"salt"`
	Salted  string    `json:"salted"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// Article returns model object for article.
type Article struct {
	ID      int64     `json:"article_id"`
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	UserID  int64     `json:"user_id"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// ArticleUser includes writer information who wrote it.
type ArticleUser struct {
	ID       int64     `json:"article_id"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	UserID   int64     `json:"user_id"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
	UserName string    `json:"user_name"`
}
