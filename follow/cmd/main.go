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

	v1 := app.Group("/api/v1/follow")
	v1.Get("/follow", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": "follow",
		})
	})

	v1.Get("/unfollow", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": "unfollow",
		})
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
