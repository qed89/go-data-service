package repositories

import (
	"database/sql"
	"go-data-service/models"
	"go-data-service/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CheckPassword(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, password FROM users WHERE username = $1`
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil // Пользователь не найден
	}
	if err != nil {
		return nil, err // Другая ошибка
	}
	return &user, nil
}

func (r *UserRepository) Save(user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRow(query, user.Username, user.Password).Scan(&user.ID)
}
