package controllers

import "github.com/gofiber/fiber/v2"

func HomePage(ctx *fiber.Ctx) error {
	return ctx.Render("home/home", nil, "layout/root")
}
