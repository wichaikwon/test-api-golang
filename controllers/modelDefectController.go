package controllers

import (
	"net/http"
	"test-api-golang/config"
	"test-api-golang/models"

	"github.com/gin-gonic/gin"
)

func GetModelDefect(c *gin.Context) {
	var modelDefects []models.ModelDefect
	config.DB.Find(&modelDefects)
	c.JSON(http.StatusOK, modelDefects)
}

func GetModelDefectById(c *gin.Context) {
	var modelDefect models.ModelDefect
	if err := config.DB.Find(&modelDefect, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, modelDefect)
}

func CreateModelDefect(c *gin.Context) {
	var modelDefect models.ModelDefect
	if err := c.ShouldBindJSON(&modelDefect); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&modelDefect)
	c.JSON(http.StatusCreated, modelDefect)
}
func UpdateModelDefect(c *gin.Context) {
	var modelDefect models.ModelDefect
	if err := config.DB.First(&modelDefect, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	if err := c.ShouldBindJSON(&modelDefect); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&modelDefect)
	c.JSON(http.StatusOK, modelDefect)
}

func DeleteModelDefect(c *gin.Context) {
	var modelDefect models.ModelDefect
	if err := config.DB.First(&modelDefect, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	config.DB.Delete(&modelDefect)
	c.JSON(http.StatusOK, gin.H{"data": "Record deleted"})
}
