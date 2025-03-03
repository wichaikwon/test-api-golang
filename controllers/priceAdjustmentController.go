package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"test-api-golang/config"
	"test-api-golang/models"
	"test-api-golang/services"

	"github.com/gin-gonic/gin"
)

func GetPriceAdjustment(c *gin.Context) {
	var adjustments []models.PriceAdjustment
	config.DB.Find(&adjustments)
	c.JSON(http.StatusOK, adjustments)
}

func GetPriceAdjustmentById(c *gin.Context) {
	var adjustment models.PriceAdjustment
	if err := config.DB.Find(&adjustment, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, adjustment)
}

func CreatePriceAdjustment(c *gin.Context) {
	var adjustment models.PriceAdjustment
	if err := c.ShouldBindJSON(&adjustment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&adjustment)
	c.JSON(http.StatusCreated, adjustment)
}

func UpdatePriceAdjustment(c *gin.Context) {
	var adjustment models.PriceAdjustment
	if err := config.DB.First(&adjustment, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	if err := c.ShouldBindJSON(&adjustment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&adjustment)
	c.JSON(http.StatusOK, adjustment)
}

func UpdateDeductions(c *gin.Context) {
	phoneID := c.Param("phone_id")

	var deductions map[string]float64
	if err := c.ShouldBindJSON(&deductions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	intDeductions := make(map[int]float64)
	for k, v := range deductions {
		intKey, err := strconv.Atoi(k)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key in deductions"})
			return
		}
		intDeductions[intKey] = v
	}

	if err := services.UpdateDeductions(phoneID, intDeductions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update deductions: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deductions updated successfully"})
}

func DeletePriceAdjustment(c *gin.Context) {
	var adjustment models.PriceAdjustment
	if err := config.DB.First(&adjustment, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	config.DB.Delete(&adjustment)
	c.JSON(http.StatusOK, gin.H{"data": "Record deleted"})
}
