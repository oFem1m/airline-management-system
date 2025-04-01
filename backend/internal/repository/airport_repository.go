package repository

import (
	"backend/internal/models"
	"database/sql"
	"errors"
	"fmt"
)

// AirportRepository определяет методы для работы с таблицей Airport
type AirportRepository struct {
	db *sql.DB
}

// NewAirportRepository возвращает новый экземпляр репозитория
func NewAirportRepository(db *sql.DB) *AirportRepository {
	return &AirportRepository{db: db}
}

// Create вставляет новую запись об аэропорте в базу данных
func (r *AirportRepository) Create(airport *models.Airport) error {
	query := `
		INSERT INTO Airports (name, code, city, country, timezone)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err := r.db.QueryRow(query,
		airport.Name,
		airport.Code,
		airport.City,
		airport.Country,
		airport.Timezone,
	).Scan(&airport.ID)

	if err != nil {
		return fmt.Errorf("не удалось создать запись в Airports: %w", err)
	}
	return nil
}

// Delete удаляет запись об аэропорте по ID
func (r *AirportRepository) Delete(id int) error {
	query := `DELETE FROM Airports WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("не удалось удалить запись из Airports: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("запись с id=%d не найдена", id)
	}

	return nil
}

// GetAll возвращает список всех аэропортов
func (r *AirportRepository) GetAll() ([]models.Airport, error) {
	query := `
		SELECT id, name, code, city, country, timezone
		FROM Airports
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("не удалось выполнить запрос GetAll: %w", err)
	}
	defer rows.Close()

	var airports []models.Airport
	for rows.Next() {
		var a models.Airport
		if err := rows.Scan(&a.ID, &a.Name, &a.Code, &a.City, &a.Country, &a.Timezone); err != nil {
			return nil, fmt.Errorf("ошибка при чтении строки: %w", err)
		}
		airports = append(airports, a)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при итерировании строк: %w", err)
	}

	return airports, nil
}

// GetByID возвращает запись аэропорта по его ID.
func (r *AirportRepository) GetByID(id int) (*models.Airport, error) {
	query := `
		SELECT id, name, code, city, country, timezone
		FROM Airports
		WHERE id = $1
	`
	var a models.Airport
	err := r.db.QueryRow(query, id).Scan(&a.ID, &a.Name, &a.Code, &a.City, &a.Country, &a.Timezone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("ошибка при получении самолёта: %w", err)
	}
	return &a, nil
}
