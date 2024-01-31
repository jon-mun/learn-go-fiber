package userDto

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CreateUserDto struct {
	Name  string `json:"name" xml:"name" form:"name" validate:"required,min=2,max=32"`
	Email string `json:"email" xml:"email" form:"email" validate:"required,email"`
}

type UpdateUserDto struct {
	Name  string `json:"name" xml:"name" form:"name" validate:"required,min=2,max=32"`
	Email string `json:"email" xml:"email" form:"email" validate:"required,email"`
}

type GetUserDto struct {
	ID    string `json:"id" xml:"id" form:"id" validate:"required,uuid"`
	Name  string `json:"name" xml:"name" form:"name" validate:"required,min=2,max=32"`
	Email string `json:"email" xml:"email" form:"email" validate:"required,email"`
}

type IError struct {
	Field string
	Tag   string
	Value string
}

var Validator = validator.New()

func ValidateCreateUser(c *fiber.Ctx) error {
	var errors []*IError
	body := new(CreateUserDto)
	c.BodyParser(&body)

	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &IError{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Param(),
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}
	return c.Next()
}

func ValidateDTO(c *fiber.Ctx, dto interface{}) error {
	var errors []*IError

	err := Validator.Struct(dto)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &IError{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Param(),
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	return c.Next()
}

func ValidationMiddleware(dtoType interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		dto := reflect.New(reflect.TypeOf(dtoType).Elem()).Interface()
		c.BodyParser(dto)
		return ValidateDTO(c, dto)
	}
}
