package handlers

import (
	"encoding/json"
	"go-data-service/models"
	"go-data-service/repositories"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	repo     *repositories.UserRepository
	validate *validator.Validate
}

func NewUserHandler(repo *repositories.UserRepository) *UserHandler {
	return &UserHandler{
		repo:     repo,
		validate: validator.New(),
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация данных
	if err := h.validate.Struct(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Сохранение пользователя
	if err := h.repo.Save(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbUser, err := h.repo.FindByUsername(user.Username)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if !h.repo.CheckPassword(user.Password, dbUser.Password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
