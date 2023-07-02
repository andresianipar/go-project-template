package router

import (
	"github.com/andresianipar/go-template/internal/router/handler"
	hv1 "github.com/andresianipar/go-template/internal/router/handler/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func (r *Router) setupRoute(app *fiber.App) {
	// Welcome
	app.Get("/", handler.Welcome)

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:          "http://example.com/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))

	v1 := app.Group("/api/v1")

	// Migration
	v1.Get("/migration", hv1.GetMigration(r.service, r.view))

	// 404 Not Found
	app.Use(handler.NotFound)
}
