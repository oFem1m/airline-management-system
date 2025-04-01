package repository

import (
	"database/sql"
	"fmt"

	"backend/internal/models"
)

// BookingRepository определяет методы для работы с таблицей Booking.
type BookingRepository struct {
	db *sql.DB
}

// NewBookingRepository возвращает новый экземпляр репозитория для бронирований.
func NewBookingRepository(db *sql.DB) *BookingRepository {
	return &BookingRepository{db: db}
}

// Create добавляет новое бронирование и возвращает сгенерированный ID.
func (r *BookingRepository) Create(booking *models.Booking) error {
	query := `
		INSERT INTO Booking (passenger_id, booking_date, status)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := r.db.QueryRow(query,
		booking.PassengerID,
		booking.BookingDate,
		booking.Status,
	).Scan(&booking.ID)
	if err != nil {
		return fmt.Errorf("не удалось создать бронирование: %w", err)
	}
	return nil
}

// Update обновляет данные бронирования по его ID.
func (r *BookingRepository) Update(booking *models.Booking) error {
	query := `
		UPDATE Booking
		SET passenger_id = $1,
		    booking_date = $2,
		    status = $3
		WHERE id = $4
	`
	result, err := r.db.Exec(query,
		booking.PassengerID,
		booking.BookingDate,
		booking.Status,
		booking.ID,
	)
	if err != nil {
		return fmt.Errorf("не удалось обновить бронирование: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("не удалось обновить бронирование: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("бронирование с id=%d не найдено", booking.ID)
	}
	return nil
}

// GetByID возвращает бронирование по его ID.
func (r *BookingRepository) GetByID(id int) (*models.Booking, error) {
	query := `
		SELECT id, passenger_id, booking_date, status
		FROM Booking
		WHERE id = $1
	`
	var booking models.Booking
	err := r.db.QueryRow(query, id).Scan(
		&booking.ID,
		&booking.PassengerID,
		&booking.BookingDate,
		&booking.Status,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("не удалось получить бронирование: %w", err)
	}
	return &booking, nil
}

// GetAll возвращает список всех бронирований.
func (r *BookingRepository) GetAll() ([]models.Booking, error) {
	query := `
		SELECT id, passenger_id, booking_date, status
		FROM Booking
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить список бронирований: %w", err)
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var b models.Booking
		if err := rows.Scan(&b.ID, &b.PassengerID, &b.BookingDate, &b.Status); err != nil {
			return nil, fmt.Errorf("ошибка при чтении бронирования: %w", err)
		}
		bookings = append(bookings, b)
	}
	return bookings, nil
}

// GetByPassenger возвращает бронирования для конкретного пассажира по его ID.
func (r *BookingRepository) GetByPassenger(passengerID int) ([]models.Booking, error) {
	query := `
		SELECT id, passenger_id, booking_date, status
		FROM Booking
		WHERE passenger_id = $1
	`
	rows, err := r.db.Query(query, passengerID)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить бронирования для пассажира: %w", err)
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var b models.Booking
		if err := rows.Scan(&b.ID, &b.PassengerID, &b.BookingDate, &b.Status); err != nil {
			return nil, fmt.Errorf("ошибка при чтении бронирования: %w", err)
		}
		bookings = append(bookings, b)
	}
	return bookings, nil
}

// Delete удаляет бронирование по его ID.
func (r *BookingRepository) Delete(id int) error {
	query := `DELETE FROM Booking WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("не удалось удалить бронирование: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("бронирование с id=%d не найдено", id)
	}
	return nil
}
