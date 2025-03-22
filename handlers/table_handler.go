package handlers

import (
	"encoding/json"
	"go-data-service/repositories"
	"net/http"
)

type TableHandler struct {
	repo *repositories.TableRepository
}

func NewTableHandler(repo *repositories.TableRepository) *TableHandler {
	return &TableHandler{repo: repo}
}

func (h *TableHandler) GetTables(w http.ResponseWriter, r *http.Request) {
	tables, err := h.repo.GetTables()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tables)
}

func (h *TableHandler) GetTableData(w http.ResponseWriter, r *http.Request) {
	tableName := r.URL.Query().Get("tableName")
	tableData, err := h.repo.GetTableData(tableName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tableData)
}
