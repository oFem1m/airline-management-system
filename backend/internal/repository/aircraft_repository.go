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
