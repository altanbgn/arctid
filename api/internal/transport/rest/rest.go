package rest

import (
  "github.com/gofiber/fiber/v2"
)

func LoadRoutes(v1 fiber.Router) {
	v1.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"status": "OK",
		})
	})
}
