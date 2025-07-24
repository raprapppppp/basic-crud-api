package db

import (
	"fmt"
	"go_fiber/config"
	"go_fiber/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the GORM database client instance
var Database *gorm.DB

func ConnectionDB() error {

	//
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Config("DB_HOST"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), config.Config("DB_PORT"))

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	Database = db
	//Database.AutoMigrate(&models.User{})
	Database.AutoMigrate(&models.Users{})
	Database.AutoMigrate(&models.Account{})
	Database.AutoMigrate(&models.Task{})

	log.Println("Database connected successfully!")
	log.Println("Database migration complete.")
	return nil

}
