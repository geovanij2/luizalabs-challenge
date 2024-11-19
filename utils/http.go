package utils

import "github.com/gofiber/fiber/v2"

type ErrorMessage struct {
	Error string `json:"message,omitempty"`
}

func RespondWithError(c *fiber.Ctx, err error, status int) error {
	return c.Status(status).JSON(ErrorMessage{
		Error: err.Error(),
	})
}

func RespondWithSuccess(c *fiber.Ctx, data interface{}, status int) error {
	return c.Status(status).JSON(data)
}

func IsAuthorized(c *fiber.Ctx, clientId string) bool {
	return c.Locals("clientId") == clientId
}
