package models

import "time"

// Booking описывает бронирование билета
type Booking struct {
	ID          int       `json:"id" db:"id"`
	PassengerID int       `json:"passenger_id" db:"passenger_id"`
	BookingDate time.Time `json:"booking_date" db:"booking_date"`
	Status      string    `json:"status" db:"status"` // например: confirmed, pending, cancelled
}
