package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"simpleng-blog-app/common/pkg/database"

	"simpleng-blog-app/follow/api"
)

func main() {
	port := os.Getenv("PORT")
	app := fiber.New()

	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")

	v1 := app.Group("/api/v1/follow")
	v1.Post("/follow", api.FollowUserHandler)

	v1.Get("/unfollow", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": "unfollow",
		})
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
