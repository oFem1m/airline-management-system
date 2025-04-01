package repository

import (
	"database/sql"
	"fmt"

	"backend/internal/models"
)

// PassengerRepository определяет методы для работы с таблицей Passengers
type PassengerRepository struct {
	db *sql.DB
}

// NewPassengerRepository возвращает новый экземпляр репозитория
func NewPassengerRepository(db *sql.DB) *PassengerRepository {
	return &PassengerRepository{db: db}
}

// Create добавляет нового пассажира в базу данных и возвращает сгенерированный ID.
func (r *PassengerRepository) Create(passenger *models.Passenger) error {
	query := `
		INSERT INTO Passengers (first_name, last_name, email, phone, passport_number)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err := r.db.QueryRow(query,
		passenger.FirstName,
		passenger.LastName,
		passenger.Email,
		passenger.Phone,
		passenger.PassportNumber,
	).Scan(&passenger.ID)
	if err != nil {
		return fmt.Errorf("не удалось создать пассажира: %w", err)
	}
	return nil
}

// Delete удаляет пассажира по его ID.
func (r *PassengerRepository) Delete(id int) error {
	query := `DELETE FROM Passengers WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("не удалось удалить пассажира: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("пассажир с id=%d не найден", id)
	}
	return nil
}

// GetAll возвращает список всех пассажиров.
func (r *PassengerRepository) GetAll() ([]models.Passenger, error) {
	query := `
		SELECT id, first_name, last_name, email, phone, passport_number
		FROM Passengers
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить список пассажиров: %w", err)
	}
	defer rows.Close()

	var passengers []models.Passenger
	for rows.Next() {
		var p models.Passenger
		if err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Email, &p.Phone, &p.PassportNumber); err != nil {
			return nil, fmt.Errorf("ошибка при чтении данных пассажира: %w", err)
		}
		passengers = append(passengers, p)
	}
	return passengers, nil
}

// GetByID возвращает данные пассажира по его ID.
func (r *PassengerRepository) GetByID(id int) (*models.Passenger, error) {
	query := `
		SELECT id, first_name, last_name, email, phone, passport_number
		FROM Passengers
		WHERE id = $1
	`
	var p models.Passenger
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.FirstName, &p.LastName, &p.Email, &p.Phone, &p.PassportNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("не удалось получить пассажира: %w", err)
	}
	return &p, nil
}

// Update обновляет данные пассажира по его ID.
func (r *PassengerRepository) Update(passenger *models.Passenger) error {
	query := `
		UPDATE Passengers
		SET first_name = $1,
		    last_name = $2,
		    email = $3,
		    phone = $4,
		    passport_number = $5
		WHERE id = $6
	`
	result, err := r.db.Exec(query,
		passenger.FirstName,
		passenger.LastName,
		passenger.Email,
		passenger.Phone,
		passenger.PassportNumber,
		passenger.ID,
	)
	if err != nil {
		return fmt.Errorf("не удалось обновить данные пассажира: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("пассажир с id=%d не найден", passenger.ID)
	}
	return nil
}
