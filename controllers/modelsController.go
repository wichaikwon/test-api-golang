package controllers

import (
	"net/http"
	"strconv"
	"test-api-golang/config"
	"test-api-golang/models"

	"github.com/gin-gonic/gin"
)

func GetModels(c *gin.Context) {
	id := c.Param("id")
	var modelsList []models.Models

	if id != "" {
		var model models.Models
		uintID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}
		if err := config.DB.First(&model, uintID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		}
		c.JSON(http.StatusOK, model)
		return
	}

	modelName := c.DefaultQuery("model_name", "")
	if modelName != "" {
		query := config.DB.Where("LOWER(model_name) LIKE LOWER(?)", "%"+modelName+"%")
		if err := query.Find(&modelsList).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
			return
		}
		c.JSON(http.StatusOK, modelsList)
		return
	}

	if err := config.DB.Find(&modelsList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.JSON(http.StatusOK, modelsList)
}

func CreateModel(c *gin.Context) {
	var model models.Models
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&model)
	c.JSON(http.StatusCreated, model)
}

func UpdateModel(c *gin.Context) {
	var model models.Models
	if err := config.DB.First(&model, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Model not found"})
		return
	}
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&model)
	c.JSON(http.StatusOK, model)
}
func DeleteModel(c *gin.Context) {
	var model models.Models
	if err := config.DB.First(&model, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Model not found"})
		return
	}
	config.DB.Delete(&model)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
