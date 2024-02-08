package main

import (
	"github.com/agilistikmal/jkt48lab-htmx/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	routes.HomeRoutes(app)
	routes.LiveRoutes(app)

	app.Static("/assets", "./assets")
	app.Get("/discord", func(c *fiber.Ctx) error {
		return c.Redirect("https://discord.gg/dVgmJfmXc2")
	})

	app.Listen(":9999")
}
