package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"backend/internal/config"
	"backend/internal/repository"
	"backend/internal/router"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	db, err := repository.ConnectDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Ошибка закрытия: %v", err)
		}
	}(db)

	r := router.NewRouter(db)

	srv := &http.Server{
		Handler:      r,
		Addr:         cfg.ServerAddress,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Запуск сервера на %s", cfg.ServerAddress)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Ошибка работы сервера: %v", err)
	}
}
