package controller

import (
	"errors"
	"log"
	"luizalabs-chalenge/application"
	"luizalabs-chalenge/domain/entity"
	"luizalabs-chalenge/utils"

	"github.com/gofiber/fiber/v2"
)

type UpdateProductController struct {
	updateProduct *application.UpdateProduct
}

func NewUpdateProductController(updateProduct *application.UpdateProduct) *UpdateProductController {
	return &UpdateProductController{
		updateProduct: updateProduct,
	}
}

func (ctrl *UpdateProductController) Handle(c *fiber.Ctx) error {
	var product entity.Product
	if err := c.BodyParser(&product); err != nil {
		log.Println(err)
		return utils.RespondWithError(c, fiber.ErrBadRequest, 400)
	}

	err := ctrl.updateProduct.Execute(&product)

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
