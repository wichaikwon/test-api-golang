package models

type PhonesView struct {
	PhoneID       int     `json:"phone_id"`
	BrandID       int     `json:"brand_id"`
	BrandName     string  `json:"brand_name"`
	ModelID       int     `json:"model_id"`
	ModelName     string  `json:"model_name"`
	CapacityID    int     `json:"capacity_id"`
	CapacityValue string  `json:"capacity_value"`
	Price         float64 `json:"price"`
}

func (PhonesView) TableName() string {
	return "view_phones"
}
