package routes

import (
	"github.com/agilistikmal/jkt48lab-htmx/controllers"
	"github.com/gofiber/fiber/v2"
)

func LiveRoutes(app *fiber.App) {
	route := app.Group("/live")

	route.Get("/", controllers.LivePage)
	route.Get("/_get-idn", controllers.LiveIDNPart)
}
