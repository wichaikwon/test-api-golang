package models

import (
	"time"
)

type Defects struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsDeleted   bool      `json:"is_deleted" gorm:"default:false"`
}
