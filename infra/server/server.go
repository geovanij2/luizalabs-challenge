package server

import (
	"luizalabs-chalenge/infra/router"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App    *fiber.App
	Router router.Router
}

func NewServer() Server {
	app := fiber.New()

	return Server{
		App:    app,
		Router: router.Router{},
	}
}

func (s Server) Run() error {
	err := s.Router.SetupRouter(s.App)

	if err != nil {
		return err
	}

	return s.App.Listen(":8080")
}
