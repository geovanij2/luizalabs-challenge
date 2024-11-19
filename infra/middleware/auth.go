package middleware

import (
	"luizalabs-chalenge/data/protocols/cryptography"
	"luizalabs-chalenge/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	decrypter cryptography.Decrypter
}

func NewAuthMiddleware(decrypter cryptography.Decrypter) *AuthMiddleware {
	return &AuthMiddleware{
		decrypter: decrypter,
	}
}

func (a *AuthMiddleware) Auth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return utils.RespondWithError(c, fiber.ErrUnauthorized, fiber.ErrUnauthorized.Code)
	}

	accessToken := authHeader[7:] // "Bearer TOKEN"

	clientId, err := a.decrypter.Decrypt(accessToken)

	if err != nil {
		return utils.RespondWithError(c, fiber.ErrUnauthorized, fiber.ErrUnauthorized.Code)
	}

	c.Locals("clientId", clientId)

	return c.Next()
}
