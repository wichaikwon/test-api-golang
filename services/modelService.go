package services

import (
	"test-api-golang/config"
	"test-api-golang/models"
)

func GetModelByID(modelID uint) (models.Models, error) {
	var model models.Models
	err := config.DB.Where("id = ?", modelID).First(&model).Error
	return model, err
}

func GetModelByName(name string) (interface{}, error) {
	var model models.Models
	result := config.DB.Where("model_name = ?", name).First(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}
