package router

import (
	"github.com/danborodin/go-chat-server/config"
	"github.com/danborodin/go-chat-server/handler"

	"github.com/gorilla/mux"
)

// Host store host address
var Host string = config.GetEnvVar("HOST")

// Router is the router variable that store all teh routes
var Router = mux.NewRouter()

// SetupRoutes set routes
func SetupRoutes() {
	// Router.Host(Host)
	// Router.Schemes("http")
	Router.HandleFunc("/login", handler.Login).Methods("POST")
	Router.HandleFunc("/register", handler.Register).Methods("POST")
	Router.HandleFunc("/channels", handler.GetChannels).Methods("GET")
	Router.HandleFunc("/channels", handler.AddNewChannel).Methods("POST")

	//Router.HandleFunc("/channel/{id}", handler.GetChannelByID).Methods("GET")
}
