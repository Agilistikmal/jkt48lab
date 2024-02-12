package routes

import (
	"github.com/agilistikmal/jkt48lab/controllers"
	"github.com/gofiber/fiber/v2"
)

func PMRoutes(app *fiber.App) {
	route := app.Group("/pm")

	route.Get("/", controllers.PMPage)
	route.Get("/_get-weekly", controllers.PMWeeklyPart)
	route.Get("/_get-monthly", controllers.PMMonthlyPart)
}
