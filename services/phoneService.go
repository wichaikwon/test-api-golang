package services

import (
	"fmt"
	"test-api-golang/config"
	"test-api-golang/models"

	"gorm.io/gorm"
)

func GetPhoneByID(phoneID uint) (models.Phones, error) {
	var phone models.Phones
	err := config.DB.Where("id = ?", phoneID).First(&phone).Error
	return phone, err
}

func GetPhonesByCriteria(db *gorm.DB, brandName, modelName, capacityValue string) ([]models.PhonesView, error) {
	var phones []models.PhonesView
	query := `
    SELECT DISTINCT ON (phone_id) * 
    FROM view_phones_with_defects 
    WHERE LOWER(brand_name) LIKE LOWER(?) 
      AND LOWER(model_name) LIKE LOWER(?) `
	err := db.Raw(query, "%"+brandName+"%", "%"+modelName).Scan(&phones).Error
	if err != nil {
		return nil, err
	}
	return phones, nil
}

func UpdateDeductions(phoneID string, deductions map[int]float64) error {
	for choiceID, deduction := range deductions {
		if err := config.DB.Model(&models.PriceAdjustment{}).
			Where("phone_id = ? AND choice_id = ?", phoneID, choiceID).
			Update("deduction", deduction).Error; err != nil {
			return fmt.Errorf("failed to update deduction for choice_id %d : %w", choiceID, err)
		}
	}
	return nil
}
