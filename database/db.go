package database

import (
	"go-api-pet/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connection() {
	if os.Getenv("env") == "DEV" {
		dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panic("Error in to connect with database")
		}
		DB.AutoMigrate(&models.Pet{})
	} else {
		DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			log.Panic("Error in to connect with database")
		}
		DB.AutoMigrate(&models.Pet{})
	}
}
