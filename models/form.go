package models

type Form struct {
	ID          string  `db:"id"`
	Name        string  `db:"name"`
	CreatedDate string  `db:"created_date"`
	UserID      int64   `db:"user_id"`
	Fields      []Field `db:"fields"` // Поля формы
}

type Field struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
}
