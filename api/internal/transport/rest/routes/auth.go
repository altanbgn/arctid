package auth

import (
  "arctid/api/internal/services"
  "github.com/gofiber/fiber/v2"
)

func Routes(v1 fiber.Router) {
  authRouter := v1.Group("/auth")

  authRouter.Post("/login", services.HandleLogin)
  authRouter.Post("/register", services.HandleRegister)
}
