package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"

	"simpleng-blog-app/users/api"
	"simpleng-blog-app/users/internal/database"
)

func main() {
	port := os.Getenv("PORT")
	app := fiber.New()

	database.Connect()
	v1 := app.Group("/api/v1/users")
	v1.Post("/register", api.RegisterUserHandler)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
