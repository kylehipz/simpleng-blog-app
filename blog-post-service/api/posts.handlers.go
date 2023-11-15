package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kylehipz/simpleng-blog-app/blog-post-service/services"
)

type PostsHandler struct {
	service *services.BlogPostService
}

func (h *PostsHandler) createBlogPost(c *fiber.Ctx) error {
	// post := h.service.CreatePost()

	return c.SendString("hello")
}

func (h *PostsHandler) getBlogPost(c *fiber.Ctx) error {
	postId := c.Params("id")

	blogPost := h.service.GetPost(postId)

	// post := h.service.GetPost(postId)

	return c.JSON(fiber.Map{
		"data": blogPost,
	})
}
