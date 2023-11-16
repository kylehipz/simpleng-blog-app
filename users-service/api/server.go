package api

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"

	"github.com/kylehipz/simpleng-blog-app/blog-post-service/services"
	"github.com/kylehipz/simpleng-blog-app/blog-post-service/store"
)

type Server struct {
	ListenAddr string
}

func (s *Server) Start() error {
	app := fiber.New()

	connStr := "postgres://postgres:postgres@blog-post-db:5432/blog-post-db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Database not connected")
	}

	log.Println("Database connected")

	postgresDB := &store.PostgresStore{DB: db}

	usersService := &services.UsersService{
		DB: postgresDB,
	}

	usersHandler := UsersHandler{
		service: usersService,
	}

	app.Get("/api/users", usersHandler.getUsers)

	log.Println("Server started on port", s.ListenAddr)
	return app.Listen(s.ListenAddr)
}

func NewServer(listenAddr string) *Server {
	return &Server{
		ListenAddr: listenAddr,
	}
}
