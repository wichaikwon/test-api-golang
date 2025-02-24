package controllers

import (
	"net/http"
	"strconv"
	"test-api-golang/services"

	"github.com/gin-gonic/gin"
)

func GetDataByIDs(c *gin.Context) {
	brandID, err1 := strconv.ParseUint(c.Param("brand_id"), 10, 32)
	modelID, err2 := strconv.ParseUint(c.Param("model_id"), 10, 32)
	phoneID, err3 := strconv.ParseUint(c.Param("phone_id"), 10, 32)

	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		c.Error(err1).SetMeta("brandID: " + c.Param("brandID"))
		return
	}
	brand, _ := services.GetBrandByID(uint(brandID))
	model, _ := services.GetModelByID(uint(modelID))
	phone, _ := services.GetPhoneByID(uint(phoneID))

	c.JSON(http.StatusOK, gin.H{"brand": brand, "model": model, "phone": phone})
}
func GetDataByNames(c *gin.Context) {
	brandName := c.Query("brand_name")
	modelName := c.Query("model_name")
	capacityName := c.Query("capacity_value")

	response := gin.H{}

	if brandName != "" {
		if brand, err := services.GetBrandByName(brandName); err == nil {
			response["brand"] = brand
		}
	}

	if modelName != "" {
		if model, err := services.GetModelByName(modelName); err == nil {
			response["model"] = model
		}
	}

	if capacityName != "" {
		if capacity, err := services.GetCapacityByName(capacityName); err == nil {
			response["capacity"] = capacity
		}
	}

	// ถ้าไม่มีข้อมูลเลยให้ส่งข้อความแจ้งเตือน
	if len(response) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No data found"})
		return
	}

	c.JSON(http.StatusOK, response)
}
