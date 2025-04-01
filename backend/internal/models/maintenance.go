package models

import "time"

// Maintenance описывает запись о техническом обслуживании
type Maintenance struct {
	ID                  int        `json:"id" db:"id"`
	AircraftID          int        `json:"aircraft_id" db:"aircraft_id"`
	MaintenanceDate     time.Time  `json:"maintenance_date" db:"maintenance_date"`
	Description         string     `json:"description" db:"description"`
	PerformedBy         *int       `json:"performed_by,omitempty" db:"performed_by"` // сотрудник, выполнивший обслуживание
	NextMaintenanceDate *time.Time `json:"next_maintenance_date,omitempty" db:"next_maintenance_date"`
}
