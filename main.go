// main.go
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jon-mun/learn-go-fiber/database"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Connect to database
	database.ConnectDB()

	// Define a route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "Hello, Fiber!",
		})
	})

	// Start the server on port 8080
	app.Listen(":8080")
}
