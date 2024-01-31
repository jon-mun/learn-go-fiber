package userHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	var responseData []interface{}
	for _, user := range users {
		responseData = append(responseData, fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		})
	}

	return c.JSON(fiber.Map{
		"data": responseData,
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

	id, err := uuid.NewUUID()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot generate UUID",
		})
	}

	user.ID = id
	result := db.Create(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB
	var user model.User

	// Convert the string ID to a UUID
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	db.Find(&user, parsedID)

	return c.JSON(fiber.Map{
		"data": user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB
	var user model.User

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	db.Find(&user, parsedID)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse User JSON",
		})
	}

	result := db.Save(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User successfully updated",
		"data": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user model.User

	parsedId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	db.First(&user, parsedId)

	if user.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No user found with given ID"})
	}

	db.Delete(&user)

	return c.SendString("User successfully deleted")
}
