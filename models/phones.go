package models

import "time"

type Phones struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	BrandID    int       `json:"brand_id" gorm:"index;references:Brands(id);onDelete:CASCADE"`
	ModelID    int       `json:"model_id" gorm:"index;references:Models(id);onDelete:CASCADE"`
	CapacityID int       `json:"capacity_id" gorm:"index;references:Capacities(id);onDelete:CASCADE"`
	Price      int       `json:"price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	IsDeleted  bool      `json:"is_deleted" gorm:"default:false"` // เปลี่ยนเป็น boolean
}
