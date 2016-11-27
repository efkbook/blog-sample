package model

import "database/sql"

// ArticlesAll returns all articles.
func ArticlesAll(db *sql.DB) ([]Article, error) {
	rows, err := db.Query(`select * from articles`)
	if err != nil {
		return nil, err
	}
	return ScanArticles(rows)
}

// ArticleOne returns the article for given id.
func ArticleOne(db *sql.DB, id int64) (Article, error) {
	return ScanArticle(db.QueryRow(`select * from articles where article_id = ?`, id))
}

func ArticleUserOne(db *sql.DB, id int64) (ArticleUser, error) {
	return ScanArticleUser(db.QueryRow(`select a.*, u.name from articles a left join users u on a.user_id = u.user_id where article_id = ?`, id))
}

// Update updates article by given article.
func (t *Article) Update(tx *sql.Tx) (sql.Result, error) {
	stmt, err := tx.Prepare(`
	update articles
		set title = ?, body = ?, updated = CURRENT_TIMESTAMP
		where article_id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(t.Title, t.Body, t.ID)
}

// Insert inserts new article.
func (t *Article) Insert(tx *sql.Tx) (sql.Result, error) {
	stmt, err := tx.Prepare(`
	insert into articles (title, body, user_id, created, updated)
	values(?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(t.Title, t.Body, t.UserID)
}

// Delete deletes article by given id.
func (t *Article) Delete(tx *sql.Tx) (sql.Result, error) {
	stmt, err := tx.Prepare(`delete from articles where article_id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(t.ID)
}

// ArticlesDeleteAll deltes all articles.
// Useful for testing.
func ArticlesDeleteAll(tx *sql.Tx) (sql.Result, error) {
	return tx.Exec(`truncate table articles`)
}
