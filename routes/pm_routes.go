package routes

import (
	"github.com/agilistikmal/jkt48lab/controllers"
	"github.com/gofiber/fiber/v2"
)

func PMRoutes(app *fiber.App) {
	route := app.Group("/pm")

	route.Get("/", controllers.PMPage)
	route.Get("/weekly", controllers.PMWeeklyPage)
	route.Get("/monthly", controllers.PMMonthlyPage)
	route.Get("/_get-weekly", controllers.PMWeeklyPart)
	route.Get("/_get-monthly", controllers.PMMonthlyPart)
	route.Get("/_get-lastweek", controllers.PMLastWeekPart)
	route.Get("/_get-lastmonth", controllers.PMLastMonthPart)
}
