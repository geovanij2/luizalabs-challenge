package controller

import (
	"errors"
	"log"
	"luizalabs-chalenge/application"
	"luizalabs-chalenge/utils"

	"github.com/gofiber/fiber/v2"
)

type ReadProductController struct {
	readProduct *application.ReadProduct
}

func NewReadProductController(readProduct *application.ReadProduct) *ReadProductController {
	return &ReadProductController{
		readProduct: readProduct,
	}
}

func (ctrl *ReadProductController) Handle(c *fiber.Ctx) error {
	var input application.ReadProductInput

	if err := c.BodyParser(&input); err != nil {
		log.Println(err)
		return utils.RespondWithError(c, fiber.ErrBadRequest, 400)
	}

	product, err := ctrl.readProduct.Execute(input)

	if err != nil && errors.Is(err, application.ErrProductNotFound) {
		log.Println(err)
		return utils.RespondWithError(c, err, 404)
	}

	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, fiber.ErrInternalServerError, 500)
	}

	return utils.RespondWithSuccess(c, product, 200)
}
