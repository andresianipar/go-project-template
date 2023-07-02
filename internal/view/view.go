package view

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type View struct {
	validate *validator.Validate
}

func New(validate *validator.Validate) *View {
	return &View{
		validate: validate,
	}
}

func (v *View) ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"data":   nil,
		"error":  err.Error(),
		"status": false,
	}
}
