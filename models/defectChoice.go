package models

import "time"

type DefectChoice struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	DefectID  int       `json:"defect_id" gorm:"index;references:Defects(id);onDelete:CASCADE"`
	Choice    string    `json:"choice"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted" gorm:"default:false"`
}
