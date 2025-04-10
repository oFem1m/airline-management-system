package models

// FlightCrew связывает рейсы и сотрудников.
type FlightCrew struct {
	FlightID   int `json:"flight_id" db:"flight_id"`
	EmployeeID int `json:"employee_id" db:"employee_id"`
}
