package repository

import (
	"database/sql"
	"fmt"

	"backend/internal/models"
)

// TicketRepository определяет методы для работы с таблицей Tickets
type TicketRepository struct {
	db *sql.DB
}

// NewTicketRepository возвращает новый экземпляр репозитория
func NewTicketRepository(db *sql.DB) *TicketRepository {
	return &TicketRepository{db: db}
}

// Create добавляет новый билет и возвращает сгенерированный ID.
func (r *TicketRepository) Create(ticket *models.Ticket) error {
	query := `
		INSERT INTO Tickets (flight_id, passenger_id, booking_id, seat_number, price, issue_date)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	err := r.db.QueryRow(query,
		ticket.FlightID,
		ticket.PassengerID,
		ticket.BookingID,
		ticket.SeatNumber,
		ticket.Price,
		ticket.IssueDate,
	).Scan(&ticket.ID)
	if err != nil {
		return fmt.Errorf("не удалось создать билет: %w", err)
	}
	return nil
}

// Update обновляет данные билета по его ID.
func (r *TicketRepository) Update(ticket *models.Ticket) error {
	query := `
		UPDATE Tickets
		SET flight_id = $1,
		    passenger_id = $2,
		    booking_id = $3,
		    seat_number = $4,
		    price = $5,
		    issue_date = $6
		WHERE id = $7
	`
	result, err := r.db.Exec(query,
		ticket.FlightID,
		ticket.PassengerID,
		ticket.BookingID,
		ticket.SeatNumber,
		ticket.Price,
		ticket.IssueDate,
		ticket.ID,
	)
	if err != nil {
		return fmt.Errorf("не удалось обновить билет: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("не удалось обновить билет: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("билет с id=%d не найден", ticket.ID)
	}
	return nil
}

// GetByID возвращает билет по его ID.
func (r *TicketRepository) GetByID(id int) (*models.Ticket, error) {
	query := `
		SELECT id, flight_id, passenger_id, booking_id, seat_number, price, issue_date
		FROM Tickets
		WHERE id = $1
	`
	var ticket models.Ticket
	err := r.db.QueryRow(query, id).Scan(
		&ticket.ID,
		&ticket.FlightID,
		&ticket.PassengerID,
		&ticket.BookingID,
		&ticket.SeatNumber,
		&ticket.Price,
		&ticket.IssueDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("не удалось получить билет: %w", err)
	}
	return &ticket, nil
}

// GetAll возвращает список всех билетов.
func (r *TicketRepository) GetAll() ([]models.Ticket, error) {
	query := `
		SELECT id, flight_id, passenger_id, booking_id, seat_number, price, issue_date
		FROM Tickets
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить список билетов: %w", err)
	}
	defer rows.Close()

	var tickets []models.Ticket
	for rows.Next() {
		var ticket models.Ticket
		if err := rows.Scan(
			&ticket.ID,
			&ticket.FlightID,
			&ticket.PassengerID,
			&ticket.BookingID,
			&ticket.SeatNumber,
			&ticket.Price,
			&ticket.IssueDate,
		); err != nil {
			return nil, fmt.Errorf("ошибка при чтении билета: %w", err)
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

// GetByPassenger возвращает билеты для конкретного пассажира по его ID.
func (r *TicketRepository) GetByPassenger(passengerID int) ([]models.Ticket, error) {
	query := `
		SELECT id, flight_id, passenger_id, booking_id, seat_number, price, issue_date
		FROM Tickets
		WHERE passenger_id = $1
	`
	rows, err := r.db.Query(query, passengerID)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить билеты для пассажира: %w", err)
	}
	defer rows.Close()

	var tickets []models.Ticket
	for rows.Next() {
		var ticket models.Ticket
		if err := rows.Scan(
			&ticket.ID,
			&ticket.FlightID,
			&ticket.PassengerID,
			&ticket.BookingID,
			&ticket.SeatNumber,
			&ticket.Price,
			&ticket.IssueDate,
		); err != nil {
			return nil, fmt.Errorf("ошибка при чтении билета: %w", err)
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

// GetByFlight возвращает билеты для конкретного рейса по его ID.
func (r *TicketRepository) GetByFlight(flightID int) ([]models.Ticket, error) {
	query := `
		SELECT id, flight_id, passenger_id, booking_id, seat_number, price, issue_date
		FROM Tickets
		WHERE flight_id = $1
	`
	rows, err := r.db.Query(query, flightID)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить билеты для рейса: %w", err)
	}
	defer rows.Close()

	var tickets []models.Ticket
	for rows.Next() {
		var ticket models.Ticket
		if err := rows.Scan(
			&ticket.ID,
			&ticket.FlightID,
			&ticket.PassengerID,
			&ticket.BookingID,
			&ticket.SeatNumber,
			&ticket.Price,
			&ticket.IssueDate,
		); err != nil {
			return nil, fmt.Errorf("ошибка при чтении билета: %w", err)
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

// GetByBooking возвращает билеты для конкретного бронирования по его ID.
func (r *TicketRepository) GetByBooking(bookingID int) ([]models.Ticket, error) {
	query := `
		SELECT id, flight_id, passenger_id, booking_id, seat_number, price, issue_date
		FROM Tickets
		WHERE booking_id = $1
	`
	rows, err := r.db.Query(query, bookingID)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить билеты для бронирования: %w", err)
	}
	defer rows.Close()

	var tickets []models.Ticket
	for rows.Next() {
		var ticket models.Ticket
		if err := rows.Scan(
			&ticket.ID,
			&ticket.FlightID,
			&ticket.PassengerID,
			&ticket.BookingID,
			&ticket.SeatNumber,
			&ticket.Price,
			&ticket.IssueDate,
		); err != nil {
			return nil, fmt.Errorf("ошибка при чтении билета: %w", err)
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

// Delete удаляет билет по его ID.
func (r *TicketRepository) Delete(id int) error {
	query := `DELETE FROM Tickets WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("не удалось удалить билет: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("билет с id=%d не найден", id)
	}
	return nil
}
