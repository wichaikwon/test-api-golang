package controllers

import (
	"net/http"
	"test-api-golang/config"
	"test-api-golang/models"

	"github.com/gin-gonic/gin"
)

func GetCapacities(c *gin.Context) {
	var capacities []models.Capacities
	config.DB.Find(&capacities)
	c.JSON(http.StatusOK, capacities)
}

func GetCapacityById(c *gin.Context) {
	var capacity models.Capacities
	if err := config.DB.First(&capacity, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, capacity)
}

func CreateCapacity(c *gin.Context) {
	var capacity models.Capacities
	if err := c.ShouldBindJSON(&capacity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&capacity)
	c.JSON(http.StatusCreated, capacity)
}

func UpdateCapacity(c *gin.Context) {
	var capacity models.Capacities
	if err := config.DB.First(&capacity, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Capacity not found"})
		return
	}
	if err := c.ShouldBindJSON(&capacity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&capacity)
	c.JSON(http.StatusOK, capacity)
}

func DeleteCapacity(c *gin.Context) {
	var capacity models.Capacities
	if err := config.DB.First(&capacity, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Capacity not found"})
		return
	}
	config.DB.Delete(&capacity)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
