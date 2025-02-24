package controllers

import (
	"net/http"
	"test-api-golang/config"
	"test-api-golang/models"

	"github.com/gin-gonic/gin"
)

func GetDefects(c *gin.Context) {
	var defects []models.Defects
	config.DB.Find(&defects)
	c.JSON(http.StatusOK, defects)
}

func GetDefectById(c *gin.Context) {
	var defect models.Defects
	if err := config.DB.Find(&defect, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, defect)
}
func CreateDefect(c *gin.Context) {
	var defect models.Defects
	if err := c.ShouldBindJSON(&defect); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&defect)
	c.JSON(http.StatusCreated, defect)
}
func UpdateDefect(c *gin.Context) {
	var defect models.Defects
	if err := config.DB.First(&defect, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	if err := c.ShouldBindJSON(&defect); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&defect)
	c.JSON(http.StatusOK, defect)
}
func DeleteDefect(c *gin.Context) {
	var defect models.Defects
	if err := config.DB.First(&defect, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	config.DB.Delete(&defect)
	c.JSON(http.StatusOK, gin.H{"data": "Record deleted"})
}
