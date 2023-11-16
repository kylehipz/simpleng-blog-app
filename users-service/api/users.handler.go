package api

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/kylehipz/simpleng-blog-app/users-service/services"
	"github.com/kylehipz/simpleng-blog-app/users-service/types"
)

type UsersHandler struct {
	service *services.UsersService
}

func (h *UsersHandler) getUsers(c *fiber.Ctx) error {
	users, err := h.service.GetUsers()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return c.JSON(fiber.Map{
		"data": users,
	})
}

func (h *UsersHandler) followUser(c *fiber.Ctx) error {
	follow := new(types.Follow)

	if err := c.BodyParser(follow); err != nil {
		log.Fatalln(err)
		return err
	}

	result, err := h.service.Follow(follow)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return c.JSON(fiber.Map{
		"data": result,
	})
}
