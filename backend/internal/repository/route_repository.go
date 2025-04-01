package repository

import (
	"database/sql"
	"fmt"

	"backend/internal/models"
)

// RouteRepository определяет методы работы с таблицей Route.
type RouteRepository struct {
	db *sql.DB
}

// NewRouteRepository возвращает новый экземпляр репозитория для маршрутов.
func NewRouteRepository(db *sql.DB) *RouteRepository {
	return &RouteRepository{db: db}
}

// Create добавляет новый маршрут в базу данных и возвращает сгенерированный ID.
func (r *RouteRepository) Create(route *models.Route) error {
	query := `
		INSERT INTO Route (departure_airport_id, arrival_airport_id, distance, duration_minutes)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	err := r.db.QueryRow(query,
		route.DepartureAirportID,
		route.ArrivalAirportID,
		route.Distance,
		route.DurationMinutes,
	).Scan(&route.ID)
	if err != nil {
		return fmt.Errorf("не удалось создать маршрут: %w", err)
	}
	return nil
}

// Delete удаляет маршрут по его ID.
func (r *RouteRepository) Delete(id int) error {
	query := `DELETE FROM Route WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("не удалось удалить маршрут: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("маршрут с id=%d не найден", id)
	}
	return nil
}

// GetAll возвращает список всех маршрутов.
func (r *RouteRepository) GetAll() ([]models.Route, error) {
	query := `
		SELECT id, departure_airport_id, arrival_airport_id, distance, duration_minutes
		FROM Route
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить список маршрутов: %w", err)
	}
	defer rows.Close()

	var routes []models.Route
	for rows.Next() {
		var rt models.Route
		if err := rows.Scan(
			&rt.ID,
			&rt.DepartureAirportID,
			&rt.ArrivalAirportID,
			&rt.Distance,
			&rt.DurationMinutes,
		); err != nil {
			return nil, fmt.Errorf("ошибка чтения строки маршрута: %w", err)
		}
		routes = append(routes, rt)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при итерации по маршрутам: %w", err)
	}
	return routes, nil
}

// GetByID возвращает маршрут по его ID.
func (r *RouteRepository) GetByID(id int) (*models.Route, error) {
	query := `
		SELECT id, departure_airport_id, arrival_airport_id, distance, duration_minutes
		FROM Route
		WHERE id = $1
	`
	var rt models.Route
	err := r.db.QueryRow(query, id).Scan(
		&rt.ID,
		&rt.DepartureAirportID,
		&rt.ArrivalAirportID,
		&rt.Distance,
		&rt.DurationMinutes,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // маршрут не найден
		}
		return nil, fmt.Errorf("ошибка получения маршрута: %w", err)
	}
	return &rt, nil
}

// GetRoutesByAirport возвращает маршруты, в которых участвует указанный аэропорт.
// Ищем как по отправлению, так и по прибытию.
func (r *RouteRepository) GetRoutesByAirport(airportID int) ([]models.Route, error) {
	query := `
		SELECT id, departure_airport_id, arrival_airport_id, distance, duration_minutes
		FROM Route
		WHERE departure_airport_id = $1 OR arrival_airport_id = $1
	`
	rows, err := r.db.Query(query, airportID)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить маршруты для аэропорта %d: %w", airportID, err)
	}
	defer rows.Close()

	var routes []models.Route
	for rows.Next() {
		var rt models.Route
		if err := rows.Scan(
			&rt.ID,
			&rt.DepartureAirportID,
			&rt.ArrivalAirportID,
			&rt.Distance,
			&rt.DurationMinutes,
		); err != nil {
			return nil, fmt.Errorf("ошибка чтения строки маршрута: %w", err)
		}
		routes = append(routes, rt)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при итерации по маршрутам: %w", err)
	}
	return routes, nil
}

// Update обновляет данные маршрута.
func (r *RouteRepository) Update(route *models.Route) error {
	query := `
		UPDATE Route
		SET departure_airport_id = $1,
		    arrival_airport_id = $2,
		    distance = $3,
		    duration_minutes = $4
		WHERE id = $5
	`
	result, err := r.db.Exec(query,
		route.DepartureAirportID,
		route.ArrivalAirportID,
		route.Distance,
		route.DurationMinutes,
		route.ID,
	)
	if err != nil {
		return fmt.Errorf("не удалось обновить маршрут: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("не удалось обновить маршрут: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("маршрут с id=%d не найден", route.ID)
	}
	return nil
}
