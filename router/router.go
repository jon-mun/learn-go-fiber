package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	userRoutes "github.com/jon-mun/learn-go-fiber/internal/routes/user"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New()) // Group endpoints with param 'api' and log whenever this endpoint is hit.

	// User routes
	userRoutes.SetupUserRoutes(api)
}
