package v1

import (
	"luizalabs-chalenge/application"
	"luizalabs-chalenge/data/protocols/cryptography"
	"luizalabs-chalenge/domain/repository"
	"luizalabs-chalenge/infra/controller"

	"github.com/gofiber/fiber/v2"
)

type LoginControllers struct {
	LoginController *controller.LoginController
}

func SetupLoginControllers(
	clientRepository repository.ClientRepository,
	encrypter cryptography.Encrypter,
	hashComparer cryptography.HashComparer,
) *LoginControllers { // usecases
	login := application.NewLogin(clientRepository, encrypter, hashComparer)
	LoginController := controller.NewLoginController(login)

	return &LoginControllers{
		LoginController: LoginController,
	}
}

func SetupV1Login(
	v1Router fiber.Router,
	clientRepository repository.ClientRepository,
	encrypter cryptography.Encrypter,
	hashComparer cryptography.HashComparer,
) {
	loginControllers := SetupLoginControllers(clientRepository, encrypter, hashComparer)

	v1Router.Post("/login", loginControllers.LoginController.Handle)
}
