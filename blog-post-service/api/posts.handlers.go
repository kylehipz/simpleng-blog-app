package api

import "github.com/gofiber/fiber/v2"

type PostsHandler struct{}

func (h *PostsHandler) createBlogPost(c *fiber.Ctx) error {
	return c.SendString("hello")
}

func (h *PostsHandler) getBlogPost(c *fiber.Ctx) error {
	postId := c.Params("id")
	return c.SendString(postId)
}
