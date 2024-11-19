package v1

import (
	"luizalabs-chalenge/infra/cryptography"
	"luizalabs-chalenge/infra/repository"

	"github.com/gofiber/fiber/v2"
)

func SetupV1Router(app *fiber.App) error {
	clientRepository := repository.NewClientRepositoryMemory()
	productRepository := repository.NewProductRepositoryMemory()
	favoritesRepository := repository.NewFavoritesRepositoryMemory()
	bcryptAdapter := cryptography.NewBcryptAdapter()
	jwtAdapter := cryptography.NewJwtAdapter("secret")

	v1Routes := app.Group("/v1")
	SetupV1Clients(v1Routes, clientRepository, productRepository, favoritesRepository, bcryptAdapter)
	SetupV1Login(v1Routes, clientRepository, jwtAdapter, bcryptAdapter)
	return nil
}
