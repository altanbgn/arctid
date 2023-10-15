package services

import (
	"arctid/api/internal/database"
	"arctid/api/internal/models"
	"arctid/api/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type LoginBody struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type RegisterBody struct {
	FirstName string `json:"firstname" form:"firstname"`
	LastName  string `json:"lastname" form:"lastname"`
	Email     string `json:"email" form:"email"`
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
}

func HandleLogin(c *fiber.Ctx) error {
	payload := new(LoginBody)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "BAD_REQUEST",
      "message": err.Error(),
		})
	}

  user := &models.User{}
  database.DB.Where("username = ?", payload.Username).First(&user)

  match, err := utils.ComparePasswordAndHash(payload.Password, user.Password)
  if err != nil || !match {
    return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
      "status": "FORBIDDEN",
      "message": "Wrong password!",
    })
  }

  token, err := utils.NewAccessToken(user.ID.String())

  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "status": "INTERNAL_SERVER_ERROR",
      "message": err.Error(),
    })
  }

  return c.JSON(fiber.Map{
    "status": "OK",
    "token": token,
  })
}

func HandleRegister(c *fiber.Ctx) error {
	payload := new(RegisterBody)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "BAD_REQUEST",
			"message": err.Error(),
		})
	}

	hashedPassword, err := utils.GenerateFromPassword(payload.Password, &utils.ArgonParams{
		Memory:      64 * 1024,
		Iterations:  1,
		Parallelism: 4,
		SaltLength:  16,
		KeyLength:   32,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "BAD_REQUEST",
			"message": err.Error(),
		})
	}

	response := database.DB.Create(&models.User{
		Firstname: payload.FirstName,
		Lastname:  payload.LastName,
		Email:     payload.Email,
		Username:  payload.Username,
		Password:  hashedPassword,
	})

	if response.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "BAD_REQUEST",
			"message": response.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "OK",
	})
}
