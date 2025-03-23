package models

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username" validate:"required,min=5,max=50"`
	Password string `db:"password" validate:"required,min=18,max=100"`
}
