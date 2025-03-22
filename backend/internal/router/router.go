package router

import (
	"database/sql"
	"net/http"

	"backend/internal/handler"
	"backend/internal/repository"

	"github.com/gorilla/mux"
)

// NewRouter настраивает маршруты приложения и возвращает http.Handler.
func NewRouter(db *sql.DB) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// Инициализируем репозиторий и обработчик для Aircraft
	aircraftRepo := repository.NewAircraftRepository(db)
	aircraftHandler := handler.NewAircraftHandler(aircraftRepo)

	// Группа маршрутов /api/v1
	api := r.PathPrefix("/api/v1").Subrouter()

	// POST /api/v1/aircraft - создание самолёта
	api.HandleFunc("/aircraft", aircraftHandler.CreateAircraft).Methods("POST")

	// GET /api/v1/aircraft – получить список всех самолётов
	api.HandleFunc("/aircrafts", aircraftHandler.GetAllAircrafts).Methods("GET")

	// DELETE /api/v1/aircraft/{id} - удаление самолёта
	api.HandleFunc("/aircraft/{id}", aircraftHandler.DeleteAircraft).Methods("DELETE")

	return r
}
