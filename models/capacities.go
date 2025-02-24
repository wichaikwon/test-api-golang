package models

import "time"

type Capacities struct {
	ID            int       `json:"id" gorm:"primaryKey;autoIncrement"`
	CapacityValue string    `json:"capacity_value"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	IsDeleted     bool      `json:"is_deleted" gorm:"default:false"`
}
