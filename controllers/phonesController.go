package controllers

import (
	"net/http"
	"test-api-golang/config"
	"test-api-golang/models"
	"test-api-golang/services"

	"github.com/gin-gonic/gin"
)

func GetPhonesByCriteria(c *gin.Context) {
	brandName := c.DefaultQuery("brand_name", "")
	modelName := c.DefaultQuery("model_name", "")
	capacityValue := c.DefaultQuery("capacity_value", "")
	phones, err := services.GetPhonesByCriteria(config.DB, brandName, modelName, capacityValue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"phones": phones})
}
func GetPhonesByBrand(c *gin.Context) {
	brand := c.Param("brand")
	model := c.DefaultQuery("model", "")
	storage := c.DefaultQuery("storage", "")
	if brand == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Brand parameter is required"})
		return
	}
	var phones []models.PhonesView
	query := config.DB.Where("brand_name = ?", brand)

	if model != "" {
		query = query.Where("model_name = ?", model)
	}
	if storage != "" {
		query = query.Where("capacity_value = ?", storage)
	}

	if err := query.Find(&phones).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	c.JSON(http.StatusOK, phones)
}

func GetPhones(c *gin.Context) {
	var phones []models.Phones
	config.DB.Find(&phones)
	c.JSON(http.StatusOK, phones)
}

func GetPhoneById(c *gin.Context) {
	var phone models.Phones
	if err := config.DB.First(&phone, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, phone)
}

func CreatePhone(c *gin.Context) {
	var phone models.Phones
	if err := c.ShouldBindJSON(&phone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&phone)
	c.JSON(http.StatusCreated, phone)
}

func UpdatePhone(c *gin.Context) {
	var phone models.Phones
	if err := config.DB.First(&phone, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Phone not found"})
		return
	}
	if err := c.ShouldBindJSON(&phone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&phone)
	c.JSON(http.StatusOK, phone)
}

func DeletePhone(c *gin.Context) {
	var phone models.Phones
	if err := config.DB.First(&phone, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Phone not found"})
		return
	}
	config.DB.Delete(&phone)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
