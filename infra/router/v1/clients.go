package v1

import (
	"luizalabs-chalenge/application"
	"luizalabs-chalenge/data/protocols/cryptography"
	"luizalabs-chalenge/domain/repository"
	"luizalabs-chalenge/infra/controller"

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
}

func SetupClientControllers(
	clientRepository repository.ClientRepository,
	productRepository repository.ProductRepository,
	favoritesRepository repository.FavoritesRepository,
	hasher cryptography.Hasher,
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

	return &ClientControllers{
		createClientController:          createClientController,
		readClientController:            readClientController,
		updateClientController:          updateClientController,
		deleteClientController:          deleteClientController,
		addFavoriteProductController:    addFavoriteProductController,
		listFavoriteProductsController:  listFavoriteProductsController,
		deleteFavoriteProductController: deleteFavoriteProductController,
	}
}

func SetupV1Clients(
	v1Router fiber.Router,
	clientRepository repository.ClientRepository,
	productRepository repository.ProductRepository,
	favoritesRepository repository.FavoritesRepository,
	hasher cryptography.Hasher,
) {
	clientControllers := SetupClientControllers(clientRepository, productRepository, favoritesRepository, hasher)

	v1Router.Post("/clients", clientControllers.createClientController.Handle)
	v1Router.Get("/clients/:clientId", clientControllers.readClientController.Handle)
	v1Router.Put("/clients/:clientId", clientControllers.updateClientController.Handle)
	v1Router.Delete("/clients/:clientId", clientControllers.deleteClientController.Handle)
	v1Router.Post("/clients/:clientId/favorites", clientControllers.addFavoriteProductController.Handle)
	v1Router.Get("/clients/:clientId/favorites", clientControllers.listFavoriteProductsController.Handle)
	v1Router.Delete("/clients/:clientId/favorites/:productId", clientControllers.deleteFavoriteProductController.Handle)
}
