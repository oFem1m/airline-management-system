package repository

import (
	"backend/internal/models"
	"database/sql"
	"fmt"
)

// AircraftRepository определяет методы для работы с таблицей Aircrafts
type AircraftRepository struct {
	db *sql.DB
}

// NewAircraftRepository возвращает новый экземпляр репозитория
func NewAircraftRepository(db *sql.DB) *AircraftRepository {
	return &AircraftRepository{db: db}
}

// Create вставляет новую запись о самолёте в базу данных
func (r *AircraftRepository) Create(aircraft *models.Aircraft) error {
	query := `
		INSERT INTO Aircrafts (tail_number, model, capacity, manufacture_year)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	err := r.db.QueryRow(query,
		aircraft.TailNumber,
		aircraft.Model,
		aircraft.Capacity,
		aircraft.ManufactureYear,
	).Scan(&aircraft.ID)

	if err != nil {
		return fmt.Errorf("не удалось создать запись в Aircrafts: %w", err)
	}
	return nil
}

// Delete удаляет запись о самолёте по ID
func (r *AircraftRepository) Delete(id int) error {
	query := `DELETE FROM Aircrafts WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("не удалось удалить запись из Aircrafts: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("запись с id=%d не найдена", id)
	}

	return nil
}

// GetAll возвращает список всех самолетов
func (r *AircraftRepository) GetAll() ([]models.Aircraft, error) {
	query := `
		SELECT id, tail_number, model, capacity, manufacture_year
		FROM Aircrafts
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("не удалось выполнить запрос GetAll: %w", err)
	}
	defer rows.Close()

	var aircrafts []models.Aircraft
	for rows.Next() {
		var a models.Aircraft
		if err := rows.Scan(&a.ID, &a.TailNumber, &a.Model, &a.Capacity, &a.ManufactureYear); err != nil {
			return nil, fmt.Errorf("ошибка при чтении строки: %w", err)
		}
		aircrafts = append(aircrafts, a)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при итерировании строк: %w", err)
	}

	return aircrafts, nil
}
