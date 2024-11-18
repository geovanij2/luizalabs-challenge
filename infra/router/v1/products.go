package v1

import (
	"luizalabs-chalenge/application"
	"luizalabs-chalenge/infra/controller"

	"github.com/gofiber/fiber/v2"
)

type ProductControllers struct {
	createProductController *controller.CreateProductController
	readProductController   *controller.ReadProductController
	updateProductController *controller.UpdateProductController
	deleteProductController *controller.DeleteProductController
}

func SetupProductControllers() *ProductControllers {
	// repositories
	productRepository := repository.NewProductRepositoryMemory()
	// usecases
	createProduct := application.NewCreateProduct(productRepository)
	readProduct := application.NewReadProduct(productRepository)
	updateProduct := application.NewUpdateProduct(productRepository)
	deleteProduct := application.NewDeleteProduct(productRepository)
	// controllers
	createProductController := controller.NewCreateProductController(createProduct)
	readProductController := controller.NewReadProductController(readProduct)
	updateProductController := controller.NewUpdateProductController(updateProduct)
	deleteProductController := controller.NewDeleteProductController(deleteProduct)

	return &ProductControllers{
		createProductController: createProductController,
		readProductController:   readProductController,
		updateProductController: updateProductController,
		deleteProductController: deleteProductController,
	}
}

func SetupV1Products(v1Router fiber.Router) {
	productControllers := SetupProductControllers()

	v1Router.Post("/products", productControllers.createProductController.Handle)
	v1Router.Get("/products/:productId", productControllers.readProductController.Handle)
	v1Router.Put("/products", productControllers.updateProductController.Handle)
	v1Router.Delete("/products/:productId", productControllers.deleteProductController.Handle)
}
