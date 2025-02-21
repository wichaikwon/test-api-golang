package models

import "time"

type Brands struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	BrandName string    `json:"brand_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted" gorm:"default:false"` // เปลี่ยนเป็น boolean
}
