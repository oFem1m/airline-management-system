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

	// Группа маршрутов /api/v1
	api := r.PathPrefix("/api/v1").Subrouter()

	// --- Самолеты ---

	// Инициализируем репозиторий и обработчик для Aircraft
	aircraftRepo := repository.NewAircraftRepository(db)
	aircraftHandler := handler.NewAircraftHandler(aircraftRepo)

	// POST - создание самолёта
	api.HandleFunc("/aircraft", aircraftHandler.CreateAircraft).Methods("POST")

	// GET – получить самолет
	api.HandleFunc("/aircraft/{id}", aircraftHandler.GetAircraft).Methods("GET")

	// GET – получить список всех самолётов
	api.HandleFunc("/aircrafts", aircraftHandler.GetAllAircrafts).Methods("GET")

	// DELETE - удаление самолёта
	api.HandleFunc("/aircraft/{id}", aircraftHandler.DeleteAircraft).Methods("DELETE")

	// --- Сотрудники ---

	// Инициализируем репозиторий и обработчик для Employee
	employeeRepo := repository.NewEmployeeRepository(db)
	employeeHandler := handler.NewEmployeeHandler(employeeRepo)

	// POST - создать сотрудника
	api.HandleFunc("/employee", employeeHandler.CreateEmployee).Methods("POST")

	// GET - список всех сотрудников
	api.HandleFunc("/employees", employeeHandler.GetAllEmployees).Methods("GET")

	// GET - получить конкретного сотрудника по ID
	api.HandleFunc("/employee/{id}", employeeHandler.GetEmployee).Methods("GET")

	// PUT - обновить данные сотрудника
	api.HandleFunc("/employee/{id}", employeeHandler.UpdateEmployee).Methods("PUT")

	// DELETE - удалить сотрудника
	api.HandleFunc("/employee/{id}", employeeHandler.DeleteEmployee).Methods("DELETE")

	// --- Аэропорты ---

	// Инициализируем репозиторий и обработчик для Airport
	airportRepo := repository.NewAirportRepository(db)
	airportHandler := handler.NewAirportHandler(airportRepo)

	// POST - создание аэропорта
	api.HandleFunc("/airport", airportHandler.CreateAirport).Methods("POST")

	// GET – получить аэропорт
	api.HandleFunc("/airport/{id}", airportHandler.GetAirport).Methods("GET")

	// GET – получить список всех аэропортов
	api.HandleFunc("/airports", airportHandler.GetAllAirports).Methods("GET")

	// DELETE - удаление аэропорта
	api.HandleFunc("/airport/{id}", airportHandler.DeleteAirport).Methods("DELETE")

	// --- Маршруты ---

	routeRepo := repository.NewRouteRepository(db)
	routeHandler := handler.NewRouteHandler(routeRepo)

	// POST – создание маршрута
	api.HandleFunc("/route", routeHandler.CreateRoute).Methods("POST")

	// DELETE – удаление маршрута по ID
	api.HandleFunc("/route/{id}", routeHandler.DeleteRoute).Methods("DELETE")

	// GET – получить список всех маршрутов
	api.HandleFunc("/routes", routeHandler.GetAllRoutes).Methods("GET")

	// GET – получить маршрут по ID
	api.HandleFunc("/route/{id}", routeHandler.GetRoute).Methods("GET")

	// GET – получить маршруты, связанные с конкретным аэропортом
	api.HandleFunc("/routes/airport/{airportId}", routeHandler.GetRoutesByAirport).Methods("GET")

	return r
}
