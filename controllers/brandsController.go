package controllers

import (
	"net/http"
	"test-api-golang/config"
	"test-api-golang/models"

	"github.com/gin-gonic/gin"
)

func GetBrands(c *gin.Context) {
	id := c.Param("id")
	var brands []models.Brands

	if id != "" {
		var brand models.Brands
		if err := config.DB.Find(&brand, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		}
		c.JSON(http.StatusOK, brand)
		return
	}
	brandName := c.DefaultQuery("brand_name", "")
	if brandName != "" {
		query := config.DB.Where("LOWER(brand_name) LIKE LOWER(?)", "%"+brandName+"%")
		if err := query.Find(&brands).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
			return
		}
		c.JSON(http.StatusOK, brands)
		return
	}
	if err := config.DB.Find(&brands).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	c.JSON(http.StatusOK, brands)
}

func CreateBrand(c *gin.Context) {
	var brand models.Brands
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&brand)
	c.JSON(http.StatusCreated, brand)
}

func UpdateBrand(c *gin.Context) {
	var brand models.Brands
	if err := config.DB.First(&brand, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&brand)
	c.JSON(http.StatusOK, brand)
}

func DeleteBrand(c *gin.Context) {
	var brand models.Brands
	if err := config.DB.First(&brand, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}
	config.DB.Delete(&brand)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
