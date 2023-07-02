package router

import (
	"github.com/andresianipar/go-template/internal/service"
	"github.com/andresianipar/go-template/internal/view"
	"github.com/gofiber/fiber/v2"
)

// @title			Nuke Deployer System API
// @version		0.1.0
// @description	Nuke Deployer System API
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.email	fiber@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:52603
// @BasePath		/
type Router struct {
	service *service.Service
	view    *view.View
}

func New(service *service.Service, view *view.View) *Router {
	return &Router{
		service: service,
		view:    view,
	}
}

func (r *Router) Setup(app *fiber.App) {
	r.setupMiddleware(app)
	r.setupRoute(app)
}
