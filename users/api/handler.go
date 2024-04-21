package api

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v3"

	"simpleng-blog-app/users/internal/usecases"
)

type RegisterRequestBody struct {
	UserName string `json:"username"`
}

func RegisterUserHandler(c fiber.Ctx) error {
	requestBody := RegisterRequestBody{}

	if err := c.Bind().Body(&requestBody); err != nil {
		log.Fatal(err)
	}

	newUser, err := usecases.RegisterUser(requestBody.UserName)
	if err != nil {
		log.Println("error occured")
		log.Printf("Error: %v\n", err)

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error occured while registering a user",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"data": newUser,
	})
}
