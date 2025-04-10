package models

import "time"

// Employee представляет сотрудника авиакомпании
type Employee struct {
	ID        int       `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	RoleID    *int      `json:"role_id,omitempty" db:"role_id"`
	HireDate  time.Time `json:"hire_date" db:"hire_date"`
	Salary    float64   `json:"salary" db:"salary"`
	Email     string    `json:"email" db:"email"`
	Phone     string    `json:"phone" db:"phone"`
}
