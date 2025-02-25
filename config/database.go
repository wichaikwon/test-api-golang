package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // for PostgreSQL driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"test-api-golang/models"
)

var DB *gorm.DB

func loadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}
}

func ConnectDB() {
	loadEnvVariables()

	// Get the DATABASE_URL from the environment
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("❌ DATABASE_URL is not set in .env file")
	}

	// Open the database connection using GORM
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	fmt.Println("✅ Database Connected Successfully!")

	// Auto-migrate models
	err = db.AutoMigrate(
		&models.Brands{},
		&models.ModelDefect{},
		&models.Models{},
		&models.Capacities{},
		&models.Phones{},
		&models.PhoneDefect{},
		&models.Defects{},
		&models.DefectChoice{},
		&models.PriceAdjustment{},
	)
	if err != nil {
		log.Fatal("❌ AutoMigrate failed:", err)
	} else {
		fmt.Println("✅ AutoMigrate completed!")
	}

	DB = db
}
