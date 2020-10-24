package main

import (
	"log"
	"net/http"

	"github.com/danborodin/go-chat-server/router"
)

func main() {
	router.SetupRoutes()
	log.Fatal(http.ListenAndServe("192.168.1.221:8000", router.Router))
}
