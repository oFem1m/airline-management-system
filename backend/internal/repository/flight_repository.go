package repository

import (
	"database/sql"
	"fmt"

	"backend/internal/models"
)

// FlightRepository определяет методы работы с таблицей Flights.
type FlightRepository struct {
	db *sql.DB
}

func NewFlightRepository(db *sql.DB) *FlightRepository {
	return &FlightRepository{db: db}
}

// Create добавляет новый рейс и возвращает сгенерированный ID.
func (r *FlightRepository) Create(flight *models.Flight) error {
	query := `
		INSERT INTO Flights (flight_number, aircraft_id, route_id, departure_time, arrival_time, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	err := r.db.QueryRow(query,
		flight.FlightNumber,
		flight.AircraftID,
		flight.RouteID,
		flight.DepartureTime,
		flight.ArrivalTime,
		flight.Status,
	).Scan(&flight.ID)
	if err != nil {
		return fmt.Errorf("не удалось создать рейс: %w", err)
	}
	return nil
}

// Delete удаляет рейс по ID.
func (r *FlightRepository) Delete(id int) error {
	query := `DELETE FROM Flights WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("не удалось удалить рейс: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("рейс с id=%d не найден", id)
	}
	return nil
}

// GetAll возвращает список всех рейсов.
func (r *FlightRepository) GetAll() ([]models.Flight, error) {
	query := `
		SELECT id, flight_number, aircraft_id, route_id, departure_time, arrival_time, status
		FROM Flights
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить список рейсов: %w", err)
	}
	defer rows.Close()

	var flights []models.Flight
	for rows.Next() {
		var f models.Flight
		if err := rows.Scan(&f.ID, &f.FlightNumber, &f.AircraftID, &f.RouteID, &f.DepartureTime, &f.ArrivalTime, &f.Status); err != nil {
			return nil, fmt.Errorf("ошибка при чтении рейса: %w", err)
		}
		flights = append(flights, f)
	}
	return flights, nil
}

// GetByID возвращает рейс по его ID.
func (r *FlightRepository) GetByID(id int) (*models.Flight, error) {
	query := `
		SELECT id, flight_number, aircraft_id, route_id, departure_time, arrival_time, status
		FROM Flights
		WHERE id = $1
	`
	var f models.Flight
	err := r.db.QueryRow(query, id).Scan(&f.ID, &f.FlightNumber, &f.AircraftID, &f.RouteID, &f.DepartureTime, &f.ArrivalTime, &f.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("ошибка получения рейса: %w", err)
	}
	return &f, nil
}

// GetByRoute возвращает все рейсы по заданному route_id.
func (r *FlightRepository) GetByRoute(routeId int) ([]models.Flight, error) {
	query := `
		SELECT id, flight_number, aircraft_id, route_id, departure_time, arrival_time, status
		FROM Flights
		WHERE route_id = $1
	`
	rows, err := r.db.Query(query, routeId)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить рейсы по маршруту: %w", err)
	}
	defer rows.Close()

	var flights []models.Flight
	for rows.Next() {
		var f models.Flight
		if err := rows.Scan(&f.ID, &f.FlightNumber, &f.AircraftID, &f.RouteID, &f.DepartureTime, &f.ArrivalTime, &f.Status); err != nil {
			return nil, fmt.Errorf("ошибка при чтении рейса: %w", err)
		}
		flights = append(flights, f)
	}
	return flights, nil
}

// GetByAircraft возвращает все рейсы для заданного aircraft_id.
func (r *FlightRepository) GetByAircraft(aircraftId int) ([]models.Flight, error) {
	query := `
		SELECT id, flight_number, aircraft_id, route_id, departure_time, arrival_time, status
		FROM Flights
		WHERE aircraft_id = $1
	`
	rows, err := r.db.Query(query, aircraftId)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить рейсы для самолёта: %w", err)
	}
	defer rows.Close()

	var flights []models.Flight
	for rows.Next() {
		var f models.Flight
		if err := rows.Scan(&f.ID, &f.FlightNumber, &f.AircraftID, &f.RouteID, &f.DepartureTime, &f.ArrivalTime, &f.Status); err != nil {
			return nil, fmt.Errorf("ошибка при чтении рейса: %w", err)
		}
		flights = append(flights, f)
	}
	return flights, nil
}

// GetByAirport возвращает рейсы для конкретного аэропорта, объединяя таблицы Flights и Route.
// Аэропорт может быть отправлением или прибитием.
func (r *FlightRepository) GetByAirport(airportId int) ([]models.Flight, error) {
	query := `
		SELECT f.id, f.flight_number, f.aircraft_id, f.route_id, f.departure_time, f.arrival_time, f.status
		FROM Flights f
		JOIN Route r ON f.route_id = r.id
		WHERE r.departure_airport_id = $1 OR r.arrival_airport_id = $1
	`
	rows, err := r.db.Query(query, airportId)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить рейсы для аэропорта: %w", err)
	}
	defer rows.Close()

	var flights []models.Flight
	for rows.Next() {
		var f models.Flight
		if err := rows.Scan(&f.ID, &f.FlightNumber, &f.AircraftID, &f.RouteID, &f.DepartureTime, &f.ArrivalTime, &f.Status); err != nil {
			return nil, fmt.Errorf("ошибка при чтении рейса: %w", err)
		}
		flights = append(flights, f)
	}
	return flights, nil
}
