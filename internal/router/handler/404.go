package handler

import "github.com/gofiber/fiber/v2"

func NotFound(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).SendString("The requested URL was not found on this server")
}
