package models

import "time"

type Element struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	HtmlCode  string    `db:"html_code"`
	CssCode   string    `db:"css_code"`
	Label     string    `db:"label"`
	UserID    int64     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}
