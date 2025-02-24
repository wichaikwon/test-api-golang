package services

import (
	"test-api-golang/config"
	"test-api-golang/models"
)

func GetCapacityByName(name string) (models.Capacities, error) {
	var capacity models.Capacities
	result := config.DB.Where("capacity_value = ?", name).First(&capacity)
	if result.Error != nil {
		return capacity, result.Error
	}
	return capacity, nil
}
