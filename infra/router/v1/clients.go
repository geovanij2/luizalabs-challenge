package v1

import (
	"luizalabs-chalenge/application"
	"luizalabs-chalenge/infra/controller"
	"luizalabs-chalenge/infra/repository"

	"github.com/gofiber/fiber/v2"
)

type ClientControllers struct {
	createClientController *controller.CreateClientController
	readClientController   *controller.ReadClientController
	updateClientController *controller.UpdateClientController
	deleteClientController *controller.DeleteClientController
}

func SetupClientControllers() *ClientControllers {
	// repositories
	clientRepository := repository.NewClientRepositoryMemory()
	// usecases
	createClient := application.NewCreateClient(clientRepository)
	readClient := application.NewReadClient(clientRepository)
	updateClient := application.NewUpdateClient(clientRepository)
	deleteClient := application.NewDeleteClient(clientRepository)
	// controllers
	createClientController := controller.NewCreateClientController(createClient)
	readClientController := controller.NewReadClientController(readClient)
	updateClientController := controller.NewUpdateClientController(updateClient)
	deleteClientController := controller.NewDeleteClientController(deleteClient)

	return &ClientControllers{
		createClientController: createClientController,
		readClientController:   readClientController,
		updateClientController: updateClientController,
		deleteClientController: deleteClientController,
	}
}

func SetupV1Clients(v1Router fiber.Router) {
	clientControllers := SetupClientControllers()

	v1Router.Post("/clients", clientControllers.createClientController.Handle)
	v1Router.Get("/clients/:clientId", clientControllers.readClientController.Handle)
	v1Router.Put("/clients", clientControllers.updateClientController.Handle)
	v1Router.Delete("/clients/:clientId", clientControllers.deleteClientController.Handle)
}
