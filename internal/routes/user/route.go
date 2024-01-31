package userRoutes

import (
	"github.com/gofiber/fiber/v2"
	userDto "github.com/jon-mun/learn-go-fiber/internal/dto/user"
	userHandler "github.com/jon-mun/learn-go-fiber/internal/handlers/user"
)

func SetupUserRoutes(router fiber.Router) {

	user := router.Group("/users")

	user.Post("/", userDto.ValidationMiddleware(new(userDto.CreateUserDto)), userHandler.CreateUser)
	user.Get("/", userHandler.GetUsers)
	user.Get("/:id", userHandler.GetUser)
	user.Put("/:id", userHandler.UpdateUser)
	user.Delete("/:id", userHandler.DeleteUser)
}
