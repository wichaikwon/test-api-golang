package services

import (
	"test-api-golang/config"
	"test-api-golang/models"
)

func GetBrandByID(brandID uint) (models.Brands, error) {
	var brand models.Brands
	err := config.DB.Where("id = ?", brandID).First(&brand).Error
	return brand, err
}

func GetBrandByName(brand_name string) (models.Brands, error) {
	var brand models.Brands
	err := config.DB.Where("brand_name = ?", brand_name).First(&brand).Error
	return brand, err
}
