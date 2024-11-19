package v1

import (
	"luizalabs-chalenge/application"
	"luizalabs-chalenge/data/protocols/cryptography"
	"luizalabs-chalenge/domain/repository"
	"luizalabs-chalenge/infra/controller"
	"luizalabs-chalenge/infra/middleware"

	"github.com/gofiber/fiber/v2"
)

type ClientControllers struct {
	createClientController          *controller.CreateClientController
	readClientController            *controller.ReadClientController
	updateClientController          *controller.UpdateClientController
	deleteClientController          *controller.DeleteClientController
	addFavoriteProductController    *controller.AddFavoriteProductController
	listFavoriteProductsController  *controller.ListFavoriteProductsController
	deleteFavoriteProductController *controller.DeleteFavoriteProductController
	authMiddleware                  *middleware.AuthMiddleware
}

func SetupClientControllers(
	clientRepository repository.ClientRepository,
	productRepository repository.ProductRepository,
	favoritesRepository repository.FavoritesRepository,
	hasher cryptography.Hasher,
	decrypter cryptography.Decrypter,
) *ClientControllers {
	// usecases
	createClient := application.NewCreateClient(clientRepository, hasher)
	readClient := application.NewReadClient(clientRepository)
	updateClient := application.NewUpdateClient(clientRepository)
	deleteClient := application.NewDeleteClient(clientRepository)
	addFavoriteProduct := application.NewAddFavoriteProduct(clientRepository, productRepository, favoritesRepository)
	listFavoriteProducts := application.NewListFavoriteProducts(favoritesRepository)
	deleteFavoriteProduct := application.NewDeleteFavoriteProduct(clientRepository, productRepository, favoritesRepository)
	// controllers
	createClientController := controller.NewCreateClientController(createClient)
	readClientController := controller.NewReadClientController(readClient)
	updateClientController := controller.NewUpdateClientController(updateClient)
	deleteClientController := controller.NewDeleteClientController(deleteClient)
	addFavoriteProductController := controller.NewAddFavoriteProductController(addFavoriteProduct)
	listFavoriteProductsController := controller.NewListFavoriteProductsController(listFavoriteProducts)
	deleteFavoriteProductController := controller.NewDeleteFavoriteProductController(deleteFavoriteProduct)
	// middleware
	authMiddleware := middleware.NewAuthMiddleware(decrypter)

	return &ClientControllers{
		createClientController:          createClientController,
		readClientController:            readClientController,
		updateClientController:          updateClientController,
		deleteClientController:          deleteClientController,
		addFavoriteProductController:    addFavoriteProductController,
		listFavoriteProductsController:  listFavoriteProductsController,
		deleteFavoriteProductController: deleteFavoriteProductController,
		authMiddleware:                  authMiddleware,
	}
}

func SetupV1Clients(
	v1Router fiber.Router,
	clientRepository repository.ClientRepository,
	productRepository repository.ProductRepository,
	favoritesRepository repository.FavoritesRepository,
	hasher cryptography.Hasher,
	decrypter cryptography.Decrypter,
) {
	clientControllers := SetupClientControllers(clientRepository, productRepository, favoritesRepository, hasher, decrypter)

	v1Router.Post("/clients", clientControllers.createClientController.Handle)
	v1Router.Get("/clients/:clientId", clientControllers.authMiddleware.Auth, clientControllers.readClientController.Handle)
	v1Router.Put("/clients/:clientId", clientControllers.authMiddleware.Auth, clientControllers.updateClientController.Handle)
	v1Router.Delete("/clients/:clientId", clientControllers.authMiddleware.Auth, clientControllers.deleteClientController.Handle)
	v1Router.Post("/clients/:clientId/favorites", clientControllers.authMiddleware.Auth, clientControllers.addFavoriteProductController.Handle)
	v1Router.Get("/clients/:clientId/favorites", clientControllers.authMiddleware.Auth, clientControllers.listFavoriteProductsController.Handle)
	v1Router.Delete("/clients/:clientId/favorites/:productId", clientControllers.authMiddleware.Auth, clientControllers.deleteFavoriteProductController.Handle)
}
