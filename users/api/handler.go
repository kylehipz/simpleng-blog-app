package api

import (
	"log"

	"github.com/gofiber/fiber/v3"

	"simpleng-blog-app/users/internal/database"
	"simpleng-blog-app/users/internal/models"
)

type RegisterRequestBody struct {
	UserName string `json:"username"`
}

func RegisterUserHandler(c fiber.Ctx) error {
	requestBody := RegisterRequestBody{}

	if err := c.Bind().Body(&requestBody); err != nil {
		log.Fatal(err)
	}

	newUser := models.User{UserName: requestBody.UserName}

	result := database.DB.Create(&newUser)

	if result.Error != nil {
		log.Println("Error occured!")
		log.Println(result.Error)

		return c.Status(400).JSON(fiber.Map{
			"message": "Attempt to create a user failed",
		})
	}

	return c.JSON(fiber.Map{
		"data": requestBody,
	})
}
