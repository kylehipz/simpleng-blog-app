package api

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v3"

	"simpleng-blog-app/blog-writer/usecases"
)

type CreateBlogRequestBody struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func CreateBlogHandler(c fiber.Ctx) error {
	requestBody := CreateBlogRequestBody{}

	if err := c.Bind().Body(&requestBody); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	newBlog, err := usecases.CreateBlog(
		requestBody.Author,
		requestBody.Title,
		requestBody.Body,
	)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error occured while creating a blog",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"data": newBlog,
	})
}

type UpdateBlogRequestBody struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func UpdateBlogHandler(c fiber.Ctx) error {
	id := c.Params("id")
	requestBody := UpdateBlogRequestBody{}

	if err := c.Bind().Body(&requestBody); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	updatedBlog, err := usecases.UpdateBlog(
		id,
		requestBody.Title,
		requestBody.Body,
	)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error occured while updating a blog",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"data": updatedBlog,
	})
}

func DeleteBlogHandler(c fiber.Ctx) error {
	id := c.Params("id")

	err := usecases.DeleteBlog(id)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error occured while deleting a blog",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}
