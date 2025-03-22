package main

import (
	"go-data-service/config"
	"go-data-service/handlers"
	"go-data-service/repositories"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	elementRepo := repositories.NewElementRepository(db)
	tableRepo := repositories.NewTableRepository(db)
	formRepo := repositories.NewFormRepository(db)

	elementHandler := handlers.NewElementHandler(elementRepo)
	tableHandler := handlers.NewTableHandler(tableRepo)
	formHandler := handlers.NewFormHandler(formRepo)

	r := mux.NewRouter()
	r.HandleFunc("/elements", elementHandler.GetElements).Methods("GET")
	r.HandleFunc("/elements/{id}", elementHandler.GetElement).Methods("GET")
	r.HandleFunc("/tables", tableHandler.GetTables).Methods("GET")
	r.HandleFunc("/tables/{tableName}", tableHandler.GetTableData).Methods("GET")
	r.HandleFunc("/forms", formHandler.GetForms).Methods("GET")
	r.HandleFunc("/forms/{id}", formHandler.GetForm).Methods("GET")
	r.HandleFunc("/forms", formHandler.SaveForm).Methods("POST")
	r.HandleFunc("/forms/{id}", formHandler.UpdateForm).Methods("PUT")
	r.HandleFunc("/forms/{id}", formHandler.DeleteForm).Methods("DELETE")

	log.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
