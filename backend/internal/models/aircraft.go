package models

// Aircraft представляет воздушное судно
type Aircraft struct {
	ID              int    `json:"id" db:"id"`
	TailNumber      string `json:"tail_number" db:"tail_number"`
	Model           string `json:"model" db:"model"`
	Capacity        int    `json:"capacity" db:"capacity"`
	ManufactureYear int    `json:"manufacture_year" db:"manufacture_year"`
}
