package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"simpleng-blog-app/common/pkg/database"

	"simpleng-blog-app/blog-writer/api"
)

func main() {
	port := os.Getenv("PORT")
	app := fiber.New()

	database.Connect()
	v1 := app.Group("/api/v1/blogs")
	v1.Post("/", api.CreateBlogHandler)
	// v1.Delete("/")
	// v1.Patch("/")

	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
