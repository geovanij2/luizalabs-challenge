package controller

import (
	"log"
	"luizalabs-chalenge/application"
	"luizalabs-chalenge/utils"

	"github.com/gofiber/fiber/v2"
)

type CreateProductController struct {
	createProduct *application.CreateProduct
}

func NewCreateProductController(createProduct *application.CreateProduct) *CreateProductController {
	return &CreateProductController{
		createProduct: createProduct,
	}
}

func (ctrl *CreateProductController) Handle(c *fiber.Ctx) error {
	var input application.CreateProductInput

	err := c.BodyParser(&input)

	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, fiber.ErrBadRequest, 400)
	}

	p, err := ctrl.createProduct.Execute(input)

	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, fiber.ErrInternalServerError, 500)
	}

	return utils.RespondWithSuccess(c, p, 201)
}
