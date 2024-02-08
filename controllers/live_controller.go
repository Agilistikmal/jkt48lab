package controllers

import (
	"github.com/agilistikmal/jkt48lab-htmx/app/service"
	"github.com/gofiber/fiber/v2"
)

func LivePage(ctx *fiber.Ctx) error {
	return ctx.Render("live/live", nil, "layout/root")
}

func LiveIDNPart(ctx *fiber.Ctx) error {
	lives := service.GetIDNLives()
	return ctx.Render("live/idn_part", fiber.Map{
		"lives": lives,
	})
}
