package main

import (
	"log"

	"github.com/danborodin/go-chat-server/database"
	"github.com/danborodin/go-chat-server/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	//db connect
	database.Connect()

	//set routes
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":6969"))

	//defer database.DB.CLose()
}
