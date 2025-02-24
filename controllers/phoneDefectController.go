package controllers

import (
	"net/http"
	"test-api-golang/config"
	"test-api-golang/models"

	"github.com/gin-gonic/gin"
)

func GetPhoneDefect(c *gin.Context) {
	var phoneDefects []models.PhoneDefect
	config.DB.Find(&phoneDefects)
	c.JSON(http.StatusOK, phoneDefects)
}
func GetPhoneDefectById(c *gin.Context) {
	var phoneDefect models.PhoneDefect
	if err := config.DB.First(&phoneDefect, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, phoneDefect)
}

func CreatePhoneDefect(c *gin.Context) {
	var phoneDefect models.PhoneDefect
	if err := c.ShouldBindJSON(&phoneDefect); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&phoneDefect)
	c.JSON(http.StatusCreated, phoneDefect)
}

func UpdatePhoneDefect(c *gin.Context) {
	var phoneDefect models.PhoneDefect
	if err := config.DB.First(&phoneDefect, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	if err := c.ShouldBindJSON(&phoneDefect); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&phoneDefect)
	c.JSON(http.StatusOK, phoneDefect)
}

func DeletePhoneDefect(c *gin.Context) {
	var phoneDefect models.PhoneDefect
	if err := config.DB.First(&phoneDefect, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	config.DB.Delete(&phoneDefect)
	c.JSON(http.StatusOK, gin.H{"data": "Record deleted"})
}
