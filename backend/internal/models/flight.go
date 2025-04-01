package models

import "time"

// Flight представляет информацию о рейсе
type Flight struct {
	ID            int       `json:"id" db:"id"`
	FlightNumber  string    `json:"flight_number" db:"flight_number"`
	AircraftID    int       `json:"aircraft_id" db:"aircraft_id"`
	RouteID       int       `json:"route_id" db:"route_id"`
	DepartureTime time.Time `json:"departure_time" db:"departure_time"`
	ArrivalTime   time.Time `json:"arrival_time" db:"arrival_time"`
	Status        string    `json:"status" db:"status"` // например: scheduled, departed, cancelled
}
