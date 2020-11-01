package main

import (
	"log"
	"net/http"

	"github.com/danborodin/go-chat-server/router"
)

func main() {
	router.SetupRoutes()
	log.Printf("Server running at address %s", router.Host)
	log.Fatal(http.ListenAndServe(router.Host, router.Router))
}
