package api

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v3"

	"simpleng-blog-app/follow/internal/usecases"
)

type FollowRequestBody struct {
	Follower string `json:"follower"`
	Followee string `json:"followee"`
}

func FollowUserHandler(c fiber.Ctx) error {
	requestBody := FollowRequestBody{}

	if err := c.Bind().Body(&requestBody); err != nil {
		log.Fatal(err)
	}

	newFollow, err := usecases.FollowUser(requestBody.Follower, requestBody.Followee)
	if err != nil {
		log.Println("error occured")
		log.Printf("Error: %v\n", err)

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error occured while following a user",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"data": newFollow,
	})
}

func UnfollowUserHandler(c fiber.Ctx) error {
	requestBody := FollowRequestBody{}

	if err := c.Bind().Body(&requestBody); err != nil {
		log.Fatal(err)
	}

	err := usecases.UnfollowUser(requestBody.Follower, requestBody.Followee)
	if err != nil {
		log.Println("error occured")
		log.Printf("Error: %v\n", err)

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error occured while following a user",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}
