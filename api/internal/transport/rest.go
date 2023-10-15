package transport

import (
  "arctid/api/internal/routes"
  "github.com/gofiber/fiber/v2"
)

func LoadRestRoutes(v1 fiber.Router) {
  auth.LoadRoutes(v1)

	v1.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"status": "OK",
		})
	})
}
