package router

import (
	v1 "luizalabs-challenge/infra/router/v1"

	"github.com/gofiber/fiber/v2"
)

type Router struct{}

func (r *Router) SetupRouter(app *fiber.App) error {
	err := v1.SetupV1Router(app)

	if err != nil {
		return err
	}

	return nil
}
