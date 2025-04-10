package repository

import (
	"database/sql"
	"fmt"

	"backend/internal/models"
)

// FlightCrewRepository определяет методы работы с таблицей FlightCrew.
type FlightCrewRepository struct {
	db *sql.DB
}

// NewFlightCrewRepository возвращает новый экземпляр репозитория для экипажей.
func NewFlightCrewRepository(db *sql.DB) *FlightCrewRepository {
	return &FlightCrewRepository{db: db}
}

// AssignEmployeeToFlight назначает сотрудника на рейс.
func (r *FlightCrewRepository) AssignEmployeeToFlight(fc *models.FlightCrew) error {
	query := `
		INSERT INTO FlightCrew (flight_id, employee_id)
		VALUES ($1, $2)
		RETURNING flight_id, employee_id
	`
	err := r.db.QueryRow(query, fc.FlightID, fc.EmployeeID).Scan(&fc.FlightID, &fc.EmployeeID)
	if err != nil {
		return fmt.Errorf("не удалось назначить сотрудника на рейс: %w", err)
	}
	return nil
}

// RemoveEmployeeFromFlight удаляет назначение сотрудника на рейс.
func (r *FlightCrewRepository) RemoveEmployeeFromFlight(flightID, employeeID int) error {
	query := `
		DELETE FROM FlightCrew
		WHERE flight_id = $1 AND employee_id = $2
	`
	result, err := r.db.Exec(query, flightID, employeeID)
	if err != nil {
		return fmt.Errorf("не удалось удалить сотрудника с рейса: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("назначение для flight_id=%d и employee_id=%d не найдено", flightID, employeeID)
	}
	return nil
}

// GetEmployeesByFlight возвращает список сотрудников, назначенных на конкретный рейс.
func (r *FlightCrewRepository) GetEmployeesByFlight(flightID int) ([]models.Employee, error) {
	query := `
		SELECT e.id, e.first_name, e.last_name, e.role_id, e.hire_date, e.salary, e.email, e.phone
		FROM FlightCrew fc
		JOIN Employees e ON fc.employee_id = e.id
		WHERE fc.flight_id = $1
	`
	rows, err := r.db.Query(query, flightID)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить сотрудников рейса: %w", err)
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var e models.Employee
		if err := rows.Scan(&e.ID, &e.FirstName, &e.LastName, &e.RoleID, &e.HireDate, &e.Salary, &e.Email, &e.Phone); err != nil {
			return nil, fmt.Errorf("ошибка при чтении данных сотрудника: %w", err)
		}
		employees = append(employees, e)
	}
	return employees, nil
}

// GetFlightsByEmployee возвращает список рейсов, на которых закреплён сотрудник.
func (r *FlightCrewRepository) GetFlightsByEmployee(employeeID int) ([]models.Flight, error) {
	query := `
		SELECT f.id, f.flight_number, f.aircraft_id, f.route_id, f.departure_time, f.arrival_time, f.status
		FROM FlightCrew fc
		JOIN Flights f ON fc.flight_id = f.id
		WHERE fc.employee_id = $1
	`
	rows, err := r.db.Query(query, employeeID)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить рейсы сотрудника: %w", err)
	}
	defer rows.Close()

	var flights []models.Flight
	for rows.Next() {
		var f models.Flight
		if err := rows.Scan(&f.ID, &f.FlightNumber, &f.AircraftID, &f.RouteID, &f.DepartureTime, &f.ArrivalTime, &f.Status); err != nil {
			return nil, fmt.Errorf("ошибка при чтении данных рейса: %w", err)
		}
		flights = append(flights, f)
	}
	return flights, nil
}
