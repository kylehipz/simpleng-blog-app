package api

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/kylehipz/simpleng-blog-app/blog-post-service/services"
)

type Server struct {
	ListenAddr string
}

func (s *Server) Start() error {
	app := fiber.New()

	postsService := &services.PostsService{
		DB: nil,
	}

	postsHandler := &PostsHandler{
		service: postsService,
	}

	app.Get("/posts/:id", postsHandler.getBlogPost)
	app.Post("/posts", postsHandler.createBlogPost)

	log.Println("Server started on port", s.ListenAddr)
	return app.Listen(s.ListenAddr)
}

func NewServer(listenAddr string) *Server {
	return &Server{
		ListenAddr: listenAddr,
	}
}
