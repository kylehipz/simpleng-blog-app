package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"simpleng-blog-app/common/pkg/database"

	"simpleng-blog-app/home-feed/internal/usecases"
)

func GetHomeFeedHandler(c fiber.Ctx) error {
	m := c.Queries()
	follower := m["follower"]
	limit, err := strconv.ParseInt(m["limit"], 10, 32)
	if err != nil {
		return err
	}

	page, err := strconv.ParseInt(m["page"], 10, 32)
	if err != nil {
		return err
	}

	count, blogs, err := usecases.GetHomeFeed(follower, int32(limit), int32(page))
	if err != nil {
		log.Println("error occured")
		log.Printf("Error: %v\n", err)

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error occured while fetching blogs",
		})
	}

	nextPage := fmt.Sprintf("/home-feed?follower=%s&limit=%d&page=%d", follower, limit, page+1)

	if len(blogs) > 0 {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"data":  blogs,
			"count": count,
			"next":  nextPage,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data":  []database.Blog{},
		"count": count,
		"next":  nextPage,
	})
}
