package handlers

import (
	"encoding/json"
	"go-data-service/models"
	"go-data-service/repositories"
	"log"
	"net/http"
	"strconv"
)

type FormHandler struct {
	repo *repositories.FormRepository
}

func NewFormHandler(repo *repositories.FormRepository) *FormHandler {
	return &FormHandler{repo: repo}
}

func (h *FormHandler) GetForms(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	log.Printf("Received GET request for forms. Page: %d, PageSize: %d\n", page, pageSize)

	forms, err := h.repo.GetForms(page, pageSize)
	if err != nil {
		log.Printf("Error fetching forms: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(forms)
}

func (h *FormHandler) GetForm(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	form, err := h.repo.GetFormByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(form)
}

func (h *FormHandler) SaveForm(w http.ResponseWriter, r *http.Request) {
	var form models.Form
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.repo.Save(&form); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *FormHandler) UpdateForm(w http.ResponseWriter, r *http.Request) {
	var form models.Form
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.repo.Update(&form); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *FormHandler) DeleteForm(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
