package controllers

import (
	"github.com/agilistikmal/jkt48lab/app/service"
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

func LiveSRPart(ctx *fiber.Ctx) error {
	lives := service.GetSRLives()
	return ctx.Render("live/sr_part", fiber.Map{
		"lives": lives,
	})
}

// IDN
func LiveIDNDetail(ctx *fiber.Ctx) error {
	return ctx.Render("live/idn/detail", fiber.Map{
		"username": ctx.Params("username"),
	}, "layout/root")
}

func LiveIDNDetailPart(ctx *fiber.Ctx) error {
	live := service.GetIDNLive(ctx.Params("username"))
	return ctx.Render("live/detail_part", fiber.Map{
		"live": live,
	})
}

func LiveIDNPlayerPart(ctx *fiber.Ctx) error {
	live := service.GetIDNLive(ctx.Params("username"))
	return ctx.Render("live/player_part", fiber.Map{
		"live": live,
	})
}

// Showroom
func LiveSRDetail(ctx *fiber.Ctx) error {
	return ctx.Render("live/sr/detail", fiber.Map{
		"username": ctx.Params("username"),
	}, "layout/root")
}

func LiveSRDetailPart(ctx *fiber.Ctx) error {
	live := service.GetSRLive(ctx.Params("username"))
	return ctx.Render("live/detail_part", fiber.Map{
		"live": live,
	})
}

func LiveSRPlayerPart(ctx *fiber.Ctx) error {
	live := service.GetSRLive(ctx.Params("username"))
	return ctx.Render("live/player_part", fiber.Map{
		"live": live,
	})
}
