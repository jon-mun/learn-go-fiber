package userHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jon-mun/learn-go-fiber/database"
	"github.com/jon-mun/learn-go-fiber/internal/model"
)

func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []model.User

	db.Find(&users)

	if len(users) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No users found",
		})
	}

	return c.JSON(fiber.Map{
		"data": users,
	})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse User JSON",
		})
	}

	db.Create(&user)

	return c.JSON(fiber.Map{
		"data": user,
	})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user model.User

	db.Find(&user, id)

	return c.JSON(fiber.Map{
		"data": user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user model.User

	db.Find(&user, id)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse User JSON",
		})
	}

	db.Save(&user)

	return c.JSON(fiber.Map{
		"data": user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user model.User

	db.First(&user, id)

	db.Delete(&user)

	return c.SendString("User successfully deleted")
}
