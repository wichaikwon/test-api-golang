package controllers

import (
	"net/http"
	"test-api-golang/config"
	"test-api-golang/models"

	"github.com/gin-gonic/gin"
)

func GetBrands(c *gin.Context) {
	var brands []models.Brands
	config.DB.Find(&brands)
	c.JSON(http.StatusOK, brands)
}
func GetBrandById(c *gin.Context) {
	var brand models.Brands
	if err := config.DB.First(&brand, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, brand)
}
func CreateBrand(c *gin.Context) {
	var input models.Brands
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&input)
	c.JSON(http.StatusOK, input)
}
