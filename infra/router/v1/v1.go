package v1

import (
	"luizalabs-chalenge/infra/cryptography"
	"luizalabs-chalenge/infra/database"
	"luizalabs-chalenge/infra/repository"

	"github.com/gofiber/fiber/v2"
)

func SetupV1Router(app *fiber.App) error {
	conn, err := database.Conn.GetConn()
	if err != nil {
		return err
	}
	clientRepository := repository.NewClientRepositoryDatabase(conn)
	productRepository := repository.NewProductRepositoryMemory()
	favoritesRepository := repository.NewFavoritesRepositoryDatabase(conn)
	bcryptAdapter := cryptography.NewBcryptAdapter()
	jwtAdapter := cryptography.NewJwtAdapter("secret")

	v1Routes := app.Group("/v1")
	SetupV1Clients(v1Routes, clientRepository, productRepository, favoritesRepository, bcryptAdapter, jwtAdapter)
	SetupV1Login(v1Routes, clientRepository, jwtAdapter, bcryptAdapter)
	return nil
}
