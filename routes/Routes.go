package routes

import (
	"test-api-golang/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func authMiddleware(c *gin.Context) {
	// Do something here
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	c.Next()

}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization"},
		AllowCredentials: true,
	}))
	r.Use(authMiddleware)

	brandRoutes := r.Group("/brands")
	{
		brandRoutes.GET("/brands", controllers.GetBrands)
		brandRoutes.GET("/brands/:id", controllers.GetBrandById)
		brandRoutes.POST("/brands", controllers.CreateBrand)
		// brandRoutes.PUT("/brands/:id", controllers.UpdateBrand)
		// brandRoutes.DELETE("/brands/:id", controllers.DeleteBrand)
	}

	return r
}
