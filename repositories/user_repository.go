package repositories

import (
	"database/sql"
	"go-data-service/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, password FROM users WHERE username = $1`
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Save(user *models.User) error {
	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRow(query, user.Username, user.Password).Scan(&user.ID)
}
