package main

import (
	"log"

	"github.com/kylehipz/simpleng-blog-app/blog-post-service/api"
)

var listenAddr = ":8000"

func main() {
	server := api.NewServer(listenAddr)

	log.Fatal(server.Start())
}
