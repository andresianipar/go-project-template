package v1

import (
	"github.com/andresianipar/go-template/internal/service"
	"github.com/andresianipar/go-template/internal/view"
	"github.com/gofiber/fiber/v2"
)

func GetMigration(service *service.Service, view *view.View) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		migrations, err := service.MigrationService().RetrieveMigrations()
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(view.ErrorResponse(err))
		}

		return ctx.JSON(view.RetrieveMigrationsResponse(migrations))
	}
}
