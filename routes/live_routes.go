package routes

import (
	"github.com/agilistikmal/jkt48lab/controllers"
	"github.com/gofiber/fiber/v2"
)

func LiveRoutes(app *fiber.App) {
	route := app.Group("/live")

	// Home
	route.Get("/", controllers.LivePage)
	route.Get("/_get-idn", controllers.LiveIDNPart)
	route.Get("/_get-sr", controllers.LiveSRPart)

	// Detail IDN
	route.Get("/idn/:username", controllers.LiveIDNDetail)
	route.Get("/idn/_get-detail/:username", controllers.LiveIDNDetailPart)
	route.Get("/idn/_get-player/:username", controllers.LiveIDNPlayerPart)

	// Detail SR
	route.Get("/sr/:username", controllers.LiveSRDetail)
	route.Get("/sr/_get-detail/:username", controllers.LiveSRDetailPart)
	route.Get("/sr/_get-player/:username", controllers.LiveSRPlayerPart)
}
