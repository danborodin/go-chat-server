package router

import (
	"github.com/danborodin/go-chat-server/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes set routes
func SetupRoutes(app *fiber.App) {

	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)
}
