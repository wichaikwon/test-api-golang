package models

import "time"

type PriceAdjustment struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	PhoneID   int       `json:"phone_id" gorm:"index;references:Phones(id);onDelete:CASCADE"`
	DefectID  int       `json:"defect_id" gorm:"index;references:Defects(id);onDelete:CASCADE"`
	ChoiceID  int       `json:"choice_id" gorm:"index;references:Choices(id);onDelete:CASCADE"`
	Deduction float64   `json:"deduction"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted" gorm:"default:false"`
}
