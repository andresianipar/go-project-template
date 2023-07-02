package handler

import (
	"fmt"

	"github.com/andresianipar/go-template/configs"
	"github.com/gofiber/fiber/v2"
)

func Welcome(ctx *fiber.Ctx) error {
	welcome := fmt.Sprintf("Welcome to the %s v%s", configs.Get("APP_NAME"), configs.Get("APP_VERSION"))
	return ctx.Status(fiber.StatusOK).Send([]byte(welcome))
}
