package router

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	// Импортируйте обработчики, когда они будут реализованы:
	// "airline-management-system/internal/handler"
)

// NewRouter настраивает маршруты приложения и возвращает http.Handler.
func NewRouter(db *sql.DB) http.Handler {
	r := mux.NewRouter()

	// Пример маршрута для проверки работоспособности
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// Пример создания группы маршрутов для публичного API (версия 1)
	// api := r.PathPrefix("/api/v1").Subrouter()
	// Пример добавления обработчика:
	// api.HandleFunc("/flights", handler.GetFlights(db)).Methods("GET")

	// Добавьте здесь другие группы маршрутов (например, для админки)

	return r
}
