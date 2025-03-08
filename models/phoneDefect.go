package models

import "time"

type PhoneDefect struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	PhoneID   int       `json:"phone_id" gorm:"index;references:Phones(id);onDelete:CASCADE"`
	DefectID  int       `json:"defect_id" gorm:"index;references:Defects(id);onDelete:CASCADE"`
	ChoiceID  int       `json:"choice_id" gorm:"index;references:DefectChoices(id);onDelete:CASCADE"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted" gorm:"default:false"`
}
