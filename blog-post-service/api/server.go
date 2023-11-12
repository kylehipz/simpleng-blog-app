package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	ListenAddr string
}

func (s *Server) Start() error {
	app := fiber.New()

	postsHandler := &PostsHandler{}

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
