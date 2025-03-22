package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	// Получаем DATABASE_URL из переменных окружения
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	// Проверяем подключение к базе данных
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database")
	return db, nil
}
