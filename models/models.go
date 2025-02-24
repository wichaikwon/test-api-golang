package models

import "time"

type Models struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	BrandID   int       `json:"brand_id" gorm:"index;references:Brands(id);onDelete:CASCADE"`
	ModelName string    `json:"model_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted" gorm:"default:false"`
}
