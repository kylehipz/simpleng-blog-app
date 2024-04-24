package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"simpleng-blog-app/common/pkg/database"

	"simpleng-blog-app/users/api"
)

func main() {
	port := os.Getenv("PORT")
	app := fiber.New()

	_, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")

	v1 := app.Group("/api/v1/users")
	v1.Post("/register", api.RegisterUserHandler)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
