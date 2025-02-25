package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// ConnectDB устанавливает соединение с базой данных по указанному URL.
func ConnectDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	// Проверяем подключение
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Успешное подключение к базе данных")
	return db, nil
}
