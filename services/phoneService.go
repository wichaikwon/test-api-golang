package services

import (
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
	query := db.Model(&models.PhonesView{}).Where("LOWER(brand_name) LIKE LOWER(?)", "%"+brandName+"%").
		Where("LOWER(model_name) LIKE LOWER(?)", "%"+modelName+"%").
		Where("LOWER(capacity_value) LIKE LOWER(?)", "%"+capacityValue+"%")

	if err := query.Find(&phones).Error; err != nil {
		return nil, err
	}

	return phones, nil
}
