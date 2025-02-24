package main

import (
	"fmt"
	"test-api-golang/config"
	"test-api-golang/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := routes.SetupRouter()
	fmt.Println("🚀 Server running on port 8080")
	r.Run(":8080")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
}
