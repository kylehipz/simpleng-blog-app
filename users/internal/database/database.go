package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"simpleng-blog-app/users/internal/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=simpleng-blog-app port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to the database")
	}

	db.AutoMigrate(&models.User{})

	fmt.Println("Successfully connected to the database")

	DB = db
}
