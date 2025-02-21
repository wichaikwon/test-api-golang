package models

import "time"

type ModelDefect struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	ModelID   int       `json:"model_id" gorm:"index;references:Models(id);onDelete:CASCADE"`
	DefectID  int       `json:"defect_id" gorm:"index;references:Defects(id);onDelete:CASCADE"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted" gorm:"default:false"` // เปลี่ยนเป็น boolean
}
