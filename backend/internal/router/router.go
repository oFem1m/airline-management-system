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

	// PUT – обновление информации об аэропорте
	api.HandleFunc("/airport/{id}", airportHandler.UpdateAirport).Methods("PUT")

	// GET – получить список всех аэропортов
	api.HandleFunc("/airports", airportHandler.GetAllAirports).Methods("GET")

	// DELETE - удаление аэропорта
	api.HandleFunc("/airport/{id}", airportHandler.DeleteAirport).Methods("DELETE")

	// --- Маршруты ---

	// Инициализируем репозиторий и обработчик для Route
	routeRepo := repository.NewRouteRepository(db)
	routeHandler := handler.NewRouteHandler(routeRepo)

	// POST – создание маршрута
	api.HandleFunc("/route", routeHandler.CreateRoute).Methods("POST")

	// PUT /api/v1/route/{id} – обновление маршрута по ID
	api.HandleFunc("/route/{id}", routeHandler.UpdateRoute).Methods("PUT")

	// DELETE – удаление маршрута по ID
	api.HandleFunc("/route/{id}", routeHandler.DeleteRoute).Methods("DELETE")

	// GET – получить список всех маршрутов
	api.HandleFunc("/routes", routeHandler.GetAllRoutes).Methods("GET")

	// GET – получить маршрут по ID
	api.HandleFunc("/route/{id}", routeHandler.GetRoute).Methods("GET")

	// GET – получить маршруты, связанные с конкретным аэропортом
	api.HandleFunc("/routes/airport/{airportId}", routeHandler.GetRoutesByAirport).Methods("GET")

	// --- Рейсы ---

	// Инициализируем репозиторий и обработчик для Flight
	flightRepo := repository.NewFlightRepository(db)
	flightHandler := handler.NewFlightHandler(flightRepo)

	// POST – создание рейса
	api.HandleFunc("/flight", flightHandler.CreateFlight).Methods("POST")

	// PUT – обновление рейса по ID
	api.HandleFunc("/flight/{id}", flightHandler.UpdateFlight).Methods("PUT")

	// DELETE – удаление рейса по ID
	api.HandleFunc("/flight/{id}", flightHandler.DeleteFlight).Methods("DELETE")

	// GET – получение всех рейсов
	api.HandleFunc("/flights", flightHandler.GetAllFlights).Methods("GET")

	// GET – получение рейса по ID
	api.HandleFunc("/flight/{id}", flightHandler.GetFlight).Methods("GET")

	// GET – получение рейсов по route_id
	api.HandleFunc("/flights/route/{routeId}", flightHandler.GetFlightsByRoute).Methods("GET")

	// GET – получение рейсов для конкретного самолёта
	api.HandleFunc("/flights/aircraft/{aircraftId}", flightHandler.GetFlightsByAircraft).Methods("GET")

	// GET – получение рейсов для конкретного аэропорта
	api.HandleFunc("/flights/airport/{airportId}", flightHandler.GetFlightsByAirport).Methods("GET")

	// --- Пассажиры ---

	// Инициализируем репозиторий и обработчик для Passenger
	passengerRepo := repository.NewPassengerRepository(db)
	passengerHandler := handler.NewPassengerHandler(passengerRepo)

	// POST – создание пассажира
	api.HandleFunc("/passenger", passengerHandler.CreatePassenger).Methods("POST")

	// GET – получение списка всех пассажиров
	api.HandleFunc("/passengers", passengerHandler.GetAllPassengers).Methods("GET")

	// GET – получение пассажира по ID
	api.HandleFunc("/passenger/{id}", passengerHandler.GetPassenger).Methods("GET")

	// PUT – обновление данных пассажира по ID
	api.HandleFunc("/passenger/{id}", passengerHandler.UpdatePassenger).Methods("PUT")

	// DELETE – удаление пассажира по ID
	api.HandleFunc("/passenger/{id}", passengerHandler.DeletePassenger).Methods("DELETE")

	// --- Бронирования ---

	// Инициализируем репозиторий и обработчик для Booking
	bookingRepo := repository.NewBookingRepository(db)
	bookingHandler := handler.NewBookingHandler(bookingRepo)

	// POST – создание бронирования
	api.HandleFunc("/booking", bookingHandler.CreateBooking).Methods("POST")

	// PUT – обновление бронирования
	api.HandleFunc("/booking/{id}", bookingHandler.UpdateBooking).Methods("PUT")

	// GET – получение бронирования по ID
	api.HandleFunc("/booking/{id}", bookingHandler.GetBooking).Methods("GET")

	// GET – получение всех бронирований
	api.HandleFunc("/bookings", bookingHandler.GetAllBookings).Methods("GET")

	// GET – получение бронирований для конкретного пассажира
	api.HandleFunc("/bookings/passenger/{passengerId}", bookingHandler.GetBookingsByPassenger).Methods("GET")

	// DELETE – удаление бронирования
	api.HandleFunc("/booking/{id}", bookingHandler.DeleteBooking).Methods("DELETE")

	// --- Билеты ---

	// Инициализируем репозиторий и обработчик для Ticket
	ticketRepo := repository.NewTicketRepository(db)
	ticketHandler := handler.NewTicketHandler(ticketRepo)

	// POST – создание билета
	api.HandleFunc("/ticket", ticketHandler.CreateTicket).Methods("POST")

	// PUT – обновление билета по ID
	api.HandleFunc("/ticket/{id}", ticketHandler.UpdateTicket).Methods("PUT")

	// GET – получение билета по ID
	api.HandleFunc("/ticket/{id}", ticketHandler.GetTicket).Methods("GET")

	// GET – получение списка всех билетов
	api.HandleFunc("/tickets", ticketHandler.GetAllTickets).Methods("GET")

	// GET – получение билетов для конкретного пассажира
	api.HandleFunc("/tickets/passenger/{passengerId}", ticketHandler.GetTicketsByPassenger).Methods("GET")

	// GET – получение билетов на рейс по ID рейса
	api.HandleFunc("/tickets/flight/{flightId}", ticketHandler.GetTicketsByFlight).Methods("GET")

	// GET – получение билетов в бронировании по ID бронирования
	api.HandleFunc("/tickets/booking/{bookingId}", ticketHandler.GetTicketsByBooking).Methods("GET")

	// DELETE – удаление билета по ID
	api.HandleFunc("/ticket/{id}", ticketHandler.DeleteTicket).Methods("DELETE")

	// --- Обслуживание ---

	// Инициализируем репозиторий и обработчик для Maintenance
	maintenanceRepo := repository.NewMaintenanceRepository(db)
	maintenanceHandler := handler.NewMaintenanceHandler(maintenanceRepo)

	// POST – создание обслуживания
	api.HandleFunc("/maintenance", maintenanceHandler.CreateMaintenance).Methods("POST")

	// PUT – обновление обслуживания по ID
	api.HandleFunc("/maintenance/{id}", maintenanceHandler.UpdateMaintenance).Methods("PUT")

	// GET – получение обслуживания по ID
	api.HandleFunc("/maintenance/{id}", maintenanceHandler.GetMaintenance).Methods("GET")

	// GET – получение списка всех обслуживаний
	api.HandleFunc("/maintenances", maintenanceHandler.GetAllMaintenance).Methods("GET")

	// GET – получение обслуживаний для конкретного самолёта
	api.HandleFunc("/maintenances/aircraft/{aircraftId}", maintenanceHandler.GetMaintenanceByAircraft).Methods("GET")

	// GET – получение обслуживаний, проведенных конкретным сотрудником
	api.HandleFunc("/maintenances/employee/{employeeId}", maintenanceHandler.GetMaintenanceByEmployee).Methods("GET")

	// DELETE – удаление обслуживания по ID
	api.HandleFunc("/maintenance/{id}", maintenanceHandler.DeleteMaintenance).Methods("DELETE")

	return r
}
