package models

// Passenger представляет данные пассажира
type Passenger struct {
	ID             int    `json:"id" db:"id"`
	FirstName      string `json:"first_name" db:"first_name"`
	LastName       string `json:"last_name" db:"last_name"`
	Email          string `json:"email" db:"email"`
	Phone          string `json:"phone" db:"phone"`
	PassportNumber string `json:"passport_number" db:"passport_number"`
}
