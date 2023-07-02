package internal

import (
	"github.com/andresianipar/go-template/internal/router"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	server *fiber.App
	router *router.Router
}

func New(
	server *fiber.App,
	router *router.Router,
) *App {
	return &App{
		server: server,
		router: router,
	}
}

func (app *App) Server() *fiber.App {
	app.router.Setup(app.server)
	return app.server
}
