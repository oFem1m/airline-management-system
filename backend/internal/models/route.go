package models

// Route описывает маршрут между двумя аэропортами
type Route struct {
	ID                 int `json:"id" db:"id"`
	DepartureAirportID int `json:"departure_airport_id" db:"departure_airport_id"`
	ArrivalAirportID   int `json:"arrival_airport_id" db:"arrival_airport_id"`
	Distance           int `json:"distance" db:"distance"`                 // расстояние в км
	DurationMinutes    int `json:"duration_minutes" db:"duration_minutes"` // длительность полёта в минутах
}
