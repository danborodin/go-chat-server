package router

import (
	"github.com/danborodin/go-chat-server/handler"

	"github.com/gorilla/mux"
)

// Host store host address
var Host string = "192.168.1.221:8000"

// Router is the router variable that store all teh routes
var Router = mux.NewRouter()

// SetupRoutes set routes
func SetupRoutes() {
	// Router.Host(Host)
	// Router.Schemes("http")
	Router.HandleFunc("/", handler.TestHandler)
	Router.HandleFunc("/login", handler.Login).Methods("POST")
	Router.HandleFunc("/register", handler.Register).Methods("POST")
}
