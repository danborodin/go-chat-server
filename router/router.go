package router

import (
	"github.com/danborodin/go-chat-server/handler"

	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

// SetupRoutes set routes
func SetupRoutes() {
	Router.HandleFunc("/", handler.TestHandler)
	Router.HandleFunc("/login", handler.Login).Methods("POST")
}
