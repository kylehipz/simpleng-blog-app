package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	port := 3000
	app := fiber.New()

	v1 := app.Group("/api/v1")
	v1.Get("/register", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": "users",
		})
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
