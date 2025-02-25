package models

import "time"

// Ticket представляет билет на рейс
type Ticket struct {
	ID          int       `json:"id" db:"id"`
	FlightID    int       `json:"flight_id" db:"flight_id"`
	PassengerID int       `json:"passenger_id" db:"passenger_id"`
	BookingID   *int      `json:"booking_id,omitempty" db:"booking_id"` // опционально
	SeatNumber  string    `json:"seat_number" db:"seat_number"`
	Price       float64   `json:"price" db:"price"`
	IssueDate   time.Time `json:"issue_date" db:"issue_date"`
}
