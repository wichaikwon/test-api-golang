package routes

import (
	"test-api-golang/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// func authMiddleware(c *gin.Context) {
// 	// Do something here
// 	token := c.GetHeader("Authorization")
// 	if token == "" {
// 		c.JSON(401, gin.H{"error": "Unauthorized"})
// 		c.Abort()
// 		return
// 	}
// 	c.Next()
// }

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// r.Use(authMiddleware)
	dataRoutes := r.Group("/api")
	{
		dataRoutes.GET("", controllers.GetPhonesByCriteria)

	}

	brandRoutes := r.Group("/")
	{
		brandRoutes.GET("/brands", controllers.GetBrands)
		brandRoutes.GET("/brands/:id", controllers.GetBrands)
		brandRoutes.POST("/brands", controllers.CreateBrand)
		brandRoutes.PUT("/brands/:id", controllers.UpdateBrand)
		brandRoutes.DELETE("/brands/:id", controllers.DeleteBrand)
	}
	capacityRoutes := r.Group("/")
	{
		capacityRoutes.GET("/capacities", controllers.GetCapacities)
		capacityRoutes.GET("/capacities/:id", controllers.GetCapacityById)
		capacityRoutes.POST("/capacities", controllers.CreateCapacity)
		capacityRoutes.PUT("/capacities/:id", controllers.UpdateCapacity)
		capacityRoutes.DELETE("/capacities/:id", controllers.DeleteCapacity)

	}

	defectChoiceRoutes := r.Group("/")
	{
		defectChoiceRoutes.GET("/defectchoice", controllers.GetDefectChoice)
		defectChoiceRoutes.GET("/defectchoice/:id", controllers.GetDefectChoiceById)
		defectChoiceRoutes.POST("/defectchoice", controllers.CreateDefectChoice)
		defectChoiceRoutes.PUT("/defectchoice/:id", controllers.UpdateDefectChoice)
		defectChoiceRoutes.DELETE("/defectchoice/:id", controllers.DeleteDefectChoice)
	}
	defectRoutes := r.Group("/")
	{
		defectRoutes.GET("/defects", controllers.GetDefects)
		defectRoutes.GET("/defects/:id", controllers.GetDefectById)
		defectRoutes.POST("/defects", controllers.CreateDefect)
		defectRoutes.PUT("/defects/:id", controllers.UpdateDefect)
		defectRoutes.DELETE("/defects/:id", controllers.DeleteDefect)
	}
	modelDefectRoutes := r.Group("/")
	{
		modelDefectRoutes.GET("/modeldefect", controllers.GetModelDefect)
		modelDefectRoutes.GET("/modeldefect/:id", controllers.GetModelDefectById)
		modelDefectRoutes.POST("/modeldefect", controllers.CreateModelDefect)
		modelDefectRoutes.PUT("/modeldefect/:id", controllers.UpdateModelDefect)
		modelDefectRoutes.DELETE("/modeldefect/:id", controllers.DeleteModelDefect)
	}
	modelRoutes := r.Group("/")
	{
		modelRoutes.GET("/models", controllers.GetModels)
		modelRoutes.POST("/models", controllers.CreateModel)
		modelRoutes.PUT("/models/:id", controllers.UpdateModel)
		modelRoutes.DELETE("/models/:id", controllers.DeleteModel)
	}
	phoneDefectRoutes := r.Group("/")
	{
		phoneDefectRoutes.GET("/phonedefect", controllers.GetPhoneDefect)
		phoneDefectRoutes.GET("/phonedefect/:id", controllers.GetPhoneDefectById)
		phoneDefectRoutes.POST("/phonedefect", controllers.CreatePhoneDefect)
		phoneDefectRoutes.PUT("/phonedefect/:id", controllers.UpdatePhoneDefect)
		phoneDefectRoutes.DELETE("/phonedefect/:id", controllers.DeletePhoneDefect)

	}

	phoneRoutes := r.Group("/")
	{
		phoneRoutes.GET("/phones", controllers.GetPhones)
		phoneRoutes.GET("/phones/:id", controllers.GetPhoneById)
		phoneRoutes.GET("/phoneswithdefects/:id", controllers.GetPhoneDetailWithDefects)
		phoneRoutes.POST("/phones", controllers.CreatePhone)
		phoneRoutes.PUT("/phones/:id", controllers.UpdatePhone)
		phoneRoutes.DELETE("/phones/:id", controllers.DeletePhone)
	}
	priceAdjusmentRoutes := r.Group("/")
	{
		priceAdjusmentRoutes.GET("/priceadjustment", controllers.GetPriceAdjustment)
		priceAdjusmentRoutes.POST("/priceadjustment", controllers.CreatePriceAdjustment)
		priceAdjusmentRoutes.PUT("/priceadjustment/:phone_id/deductions", controllers.UpdateDeductions)

	}

	return r
}
