package routes

import (
	"github.com/agilistikmal/jkt48lab-htmx/controllers"
	"github.com/gofiber/fiber/v2"
)

func HomeRoutes(app *fiber.App) {
	route := app.Group("/")

	route.Get("/", controllers.HomePage)
}
