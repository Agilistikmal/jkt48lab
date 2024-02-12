package controllers

import (
	"github.com/agilistikmal/jkt48lab/app/service"
	"github.com/gofiber/fiber/v2"
)

func PMPage(ctx *fiber.Ctx) error {
	return ctx.Render("pm/pm", nil, "layout/root")
}

func PMWeeklyPart(ctx *fiber.Ctx) error {
	stats := service.GetPMStats("weekly")
	return ctx.Render("pm/pm_part", fiber.Map{
		"stats": stats,
	})
}

func PMMonthlyPart(ctx *fiber.Ctx) error {
	stats := service.GetPMStats("monthly")
	return ctx.Render("pm/pm_part", fiber.Map{
		"stats": stats,
	})
}
