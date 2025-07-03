package db

import (
	"go_fiber/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the GORM database client instance
var Database *gorm.DB

func ConnectionDB(dsn string) error {
	var err error
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	Database.AutoMigrate(&models.User{})
	log.Println("Database connected successfully!")
	log.Println("Database migration complete.")
	return nil

}
