package repository

import (
	"database/sql"
	"fmt"

	"backend/internal/models"
)

// MaintenanceRepository определяет методы для работы с таблицей Maintenance
type MaintenanceRepository struct {
	db *sql.DB
}

// NewMaintenanceRepository возвращает новый экземпляр репозитория
func NewMaintenanceRepository(db *sql.DB) *MaintenanceRepository {
	return &MaintenanceRepository{db: db}
}

// Create добавляет новую запись обслуживания и возвращает сгенерированный ID.
func (r *MaintenanceRepository) Create(m *models.Maintenance) error {
	query := `
		INSERT INTO Maintenance (aircraft_id, maintenance_date, description, performed_by, next_maintenance_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err := r.db.QueryRow(query,
		m.AircraftID,
		m.MaintenanceDate,
		m.Description,
		m.PerformedBy,
		m.NextMaintenanceDate,
	).Scan(&m.ID)
	if err != nil {
		return fmt.Errorf("не удалось создать запись обслуживания: %w", err)
	}
	return nil
}

// Update обновляет данные обслуживания по его ID.
func (r *MaintenanceRepository) Update(m *models.Maintenance) error {
	query := `
		UPDATE Maintenance
		SET aircraft_id = $1,
		    maintenance_date = $2,
		    description = $3,
		    performed_by = $4,
		    next_maintenance_date = $5
		WHERE id = $6
	`
	result, err := r.db.Exec(query,
		m.AircraftID,
		m.MaintenanceDate,
		m.Description,
		m.PerformedBy,
		m.NextMaintenanceDate,
		m.ID,
	)
	if err != nil {
		return fmt.Errorf("не удалось обновить обслуживание: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("не удалось обновить обслуживание: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("обслуживание с id=%d не найдено", m.ID)
	}
	return nil
}

// GetByID возвращает запись обслуживания по его ID.
func (r *MaintenanceRepository) GetByID(id int) (*models.Maintenance, error) {
	query := `
		SELECT id, aircraft_id, maintenance_date, description, performed_by, next_maintenance_date
		FROM Maintenance
		WHERE id = $1
	`
	var m models.Maintenance
	err := r.db.QueryRow(query, id).Scan(
		&m.ID,
		&m.AircraftID,
		&m.MaintenanceDate,
		&m.Description,
		&m.PerformedBy,
		&m.NextMaintenanceDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("не удалось получить обслуживание: %w", err)
	}
	return &m, nil
}

// GetAll возвращает список всех записей обслуживания.
func (r *MaintenanceRepository) GetAll() ([]models.Maintenance, error) {
	query := `
		SELECT id, aircraft_id, maintenance_date, description, performed_by, next_maintenance_date
		FROM Maintenance
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить список обслуживаний: %w", err)
	}
	defer rows.Close()

	var maintenances []models.Maintenance
	for rows.Next() {
		var m models.Maintenance
		if err := rows.Scan(
			&m.ID,
			&m.AircraftID,
			&m.MaintenanceDate,
			&m.Description,
			&m.PerformedBy,
			&m.NextMaintenanceDate,
		); err != nil {
			return nil, fmt.Errorf("ошибка при чтении обслуживания: %w", err)
		}
		maintenances = append(maintenances, m)
	}
	return maintenances, nil
}

// GetByAircraft возвращает все обслуживания для конкретного самолёта по его ID.
func (r *MaintenanceRepository) GetByAircraft(aircraftID int) ([]models.Maintenance, error) {
	query := `
		SELECT id, aircraft_id, maintenance_date, description, performed_by, next_maintenance_date
		FROM Maintenance
		WHERE aircraft_id = $1
	`
	rows, err := r.db.Query(query, aircraftID)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить обслуживания для самолёта: %w", err)
	}
	defer rows.Close()

	var maintenances []models.Maintenance
	for rows.Next() {
		var m models.Maintenance
		if err := rows.Scan(
			&m.ID,
			&m.AircraftID,
			&m.MaintenanceDate,
			&m.Description,
			&m.PerformedBy,
			&m.NextMaintenanceDate,
		); err != nil {
			return nil, fmt.Errorf("ошибка при чтении обслуживания: %w", err)
		}
		maintenances = append(maintenances, m)
	}
	return maintenances, nil
}

// GetByEmployee возвращает все обслуживания, проведённые конкретным сотрудником по его ID.
func (r *MaintenanceRepository) GetByEmployee(employeeID int) ([]models.Maintenance, error) {
	query := `
		SELECT id, aircraft_id, maintenance_date, description, performed_by, next_maintenance_date
		FROM Maintenance
		WHERE performed_by = $1
	`
	rows, err := r.db.Query(query, employeeID)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить обслуживания для сотрудника: %w", err)
	}
	defer rows.Close()

	var maintenances []models.Maintenance
	for rows.Next() {
		var m models.Maintenance
		if err := rows.Scan(
			&m.ID,
			&m.AircraftID,
			&m.MaintenanceDate,
			&m.Description,
			&m.PerformedBy,
			&m.NextMaintenanceDate,
		); err != nil {
			return nil, fmt.Errorf("ошибка при чтении обслуживания: %w", err)
		}
		maintenances = append(maintenances, m)
	}
	return maintenances, nil
}

// Delete удаляет запись обслуживания по его ID.
func (r *MaintenanceRepository) Delete(id int) error {
	query := `DELETE FROM Maintenance WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("не удалось удалить обслуживание: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("обслуживание с id=%d не найдено", id)
	}
	return nil
}
