package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
)

func main() {
	port := os.Getenv("PORT")
	app := fiber.New()

	v1 := app.Group("/api/v1/blog")
	v1.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": "users",
		})
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
