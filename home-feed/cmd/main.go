package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"simpleng-blog-app/common/pkg/database"

	"simpleng-blog-app/home-feed/api"
)

func main() {
	port := os.Getenv("PORT")
	app := fiber.New()

	err := database.Connect()
	if err != nil {
		log.Fatal("Error connecting to the database")
	}

	log.Println("Database connected")

	v1 := app.Group("/api/v1/home-feed")
	v1.Get("/", api.GetHomeFeedHandler)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
