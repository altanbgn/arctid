package services

import (
	// "log"

  // "arctid/api/pkg/database"

	"github.com/gofiber/fiber/v2"
)

type LoginBody struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type RegisterBody struct {
  FirstName string `json:"firstname" form:"firstname"`
  LastName string `json:"lastname" form:"lastname"`
  Email string `json:"email" form:"email"`
  Username string `json:"username" form:"username"`
  Password string `json:"password" form:"password"`
}

func HandleLogin(c *fiber.Ctx) error {
	payload := new(LoginBody)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "BAD_REQUEST",
		})
	}

	return c.JSON(fiber.Map{
		"status": "OK",
	})
}

func HandleRegister(c *fiber.Ctx) error {
  payload := new(RegisterBody)

  if err := c.BodyParser(payload); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "status": "BAD_REQUEST",
    })
  }

	return c.JSON(fiber.Map{
		"status": "OK",
	})
}
