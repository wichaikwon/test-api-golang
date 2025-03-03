package models

type PhonesView struct {
	PhoneID        int     `json:"phone_id" gorm:"column:phone_id"`
	BrandID        int     `json:"brand_id" gorm:"column:brand_id"`
	BrandName      string  `json:"brand_name" gorm:"column:brand_name"`
	ModelID        int     `json:"model_id" gorm:"column:model_id"`
	ModelName      string  `json:"model_name" gorm:"column:model_name"`
	CapacityID     int     `json:"capacity_id" gorm:"column:capacity_id"`
	CapacityValue  string  `json:"capacity_value" gorm:"column:capacity_value"`
	Price          float64 `json:"price" gorm:"column:price"`
	MinPrice       float64 `json:"min_price" gorm:"column:min_price"`
	DefectID       int     `json:"defect_id" gorm:"column:defect_id"`
	DefectCategory string  `json:"defect_category" gorm:"column:defect_category"`
	DefectDesc     string  `json:"defect_desc" gorm:"column:defect_description"`
	ChoiceID       int     `json:"choice_id" gorm:"column:choice_id"`
	DefectChoice   string  `json:"defect_choice" gorm:"column:defect_choice"`
	Deduction      float64 `json:"deduction" gorm:"column:deduction"`
}

func (PhonesView) TableName() string {
	return "view_phones_with_defects"
}
