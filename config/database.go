package config

import (
	"byfood-test-backend/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
}

func MigrateDatabase() {

	DB.AutoMigrate(&models.Book{})

}

func SetupTestDB() {

	println("HELLO THERE", os.Getenv("TEST_DB_URL"))

	var err error
	dsn := os.Getenv("TEST_DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	DB.AutoMigrate(&models.Book{})
}
