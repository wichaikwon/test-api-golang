package controllers

import (
	"net/http"
	"strconv"
	"strings"
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

func GetPhones(c *gin.Context) {
	var phones []models.Phones
	config.DB.Find(&phones)
	c.JSON(http.StatusOK, phones)
}
func GetPhoneDetailWithDefects(c *gin.Context) {
	phoneID := c.Param("id")
	var phones []models.PhonesView
	if err := config.DB.Where("phone_id = ?", phoneID).Order("defect_id, choice_id ASC").Find(&phones).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to retrieve phone details",
		})
		return
	}
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

func GetFinalPrice(c *gin.Context) {
	phoneID := c.Query("phone_id")
	choiceIDs := c.Query("choice_id")

	phoneIDInt, err := strconv.Atoi(phoneID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid phone_id"})
		return
	}

	choiceIDStrs := strings.Split(choiceIDs, ",")
	var choiceIDInts []int
	for _, id := range choiceIDStrs {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid choice_id"})
			return
		}
		choiceIDInts = append(choiceIDInts, idInt)
	}
	finalPrice, err := models.CalculateFinalPrice(config.DB, phoneIDInt, choiceIDInts)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"final_price": finalPrice})
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
