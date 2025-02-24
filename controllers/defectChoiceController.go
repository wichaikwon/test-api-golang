package controllers

import (
	"net/http"
	"test-api-golang/config"
	"test-api-golang/models"

	"github.com/gin-gonic/gin"
)

func GetDefectChoice(c *gin.Context) {
	var defectChoices []models.DefectChoice
	config.DB.Find(&defectChoices)
	c.JSON(http.StatusOK, defectChoices)
}

func GetDefectChoiceById(c *gin.Context) {
	var defectChoice models.DefectChoice
	if err := config.DB.Find(&defectChoice, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, defectChoice)
}

func CreateDefectChoice(c *gin.Context) {
	var defectChoice models.DefectChoice
	if err := c.ShouldBindJSON(&defectChoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&defectChoice)
	c.JSON(http.StatusCreated, defectChoice)
}

func UpdateDefectChoice(c *gin.Context) {
	var defectChoice models.DefectChoice
	if err := config.DB.First(&defectChoice, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	if err := c.ShouldBindJSON(&defectChoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&defectChoice)
	c.JSON(http.StatusOK, defectChoice)
}

func DeleteDefectChoice(c *gin.Context) {
	var defectChoice models.DefectChoice
	if err := config.DB.First(&defectChoice, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	config.DB.Delete(&defectChoice)
	c.JSON(http.StatusOK, gin.H{"data": "Record deleted"})
}
