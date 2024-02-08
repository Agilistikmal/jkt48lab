package controllers

import "github.com/gofiber/fiber/v2"

func HomePage(ctx *fiber.Ctx) error {
	return ctx.Render("home/home", fiber.Map{
		"text": "Hello World",
	}, "layout/root")
}
