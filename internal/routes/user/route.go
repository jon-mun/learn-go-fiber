package userRoutes

import (
	"github.com/gofiber/fiber/v2"
	userHandler "github.com/jon-mun/learn-go-fiber/internal/handlers/user"
)

func SetupUserRoutes(router fiber.Router) {

	user := router.Group("/users")

	user.Post("/", userHandler.CreateUser)
	user.Get("/", userHandler.GetUsers)
	user.Get("/:id", userHandler.GetUser)
	user.Put("/:id", userHandler.UpdateUser)
	user.Delete("/:id", userHandler.DeleteUser)
}
