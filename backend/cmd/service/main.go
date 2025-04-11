package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"backend/internal/config"
	"backend/internal/repository"
	"backend/internal/router"
	"github.com/gorilla/handlers"
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

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	srv := &http.Server{
		Handler:      corsHandler,
		Addr:         cfg.ServerAddress,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Запуск сервера на %s", cfg.ServerAddress)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Ошибка работы сервера: %v", err)
	}
}
