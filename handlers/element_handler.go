package handlers

import (
	"encoding/json"
	"go-data-service/repositories"
	"net/http"
	"strconv"
)

type ElementHandler struct {
	repo *repositories.ElementRepository
}

func NewElementHandler(repo *repositories.ElementRepository) *ElementHandler {
	return &ElementHandler{repo: repo}
}

func (h *ElementHandler) GetElements(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	elements, err := h.repo.GetElements(page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(elements)
}

func (h *ElementHandler) GetElement(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	element, err := h.repo.GetElementByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(element)
}
