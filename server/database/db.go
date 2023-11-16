package db

import (
	"log"
	"os"

	"github.com/abdulmanafc2001/microservice-api-usercrud/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// creating global variable for accessing main package

func ConnectToDB() *gorm.DB {
	//geting data from env file
	dsn := os.Getenv("DB_URL")

	// connecting to postgres database
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	log.Println("connected to database")

	// creating a table
	DB.AutoMigrate(&models.User{})

	return DB
}
