package models

// Airport представляет данные об аэропорте
type Airport struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Code     string `json:"code" db:"code"`
	City     string `json:"city" db:"city"`
	Country  string `json:"country" db:"country"`
	Timezone string `json:"timezone" db:"timezone"`
}
