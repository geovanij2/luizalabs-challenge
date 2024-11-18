package controller

import (
	"errors"
	"luizalabs-chalenge/application"
	"luizalabs-chalenge/utils"

	"github.com/gofiber/fiber/v2"
)

type DeleteProductController struct {
	deleteProduct *application.DeleteProduct
}

func NewDeleteProductController(deleteProduct *application.DeleteProduct) *DeleteProductController {
	return &DeleteProductController{
		deleteProduct: deleteProduct,
	}
}

func (ctrl *DeleteProductController) Handle(c *fiber.Ctx) error {
	productId := c.Params("productId")

	if productId == "" {
		return utils.RespondWithError(c, fiber.ErrBadRequest, 400)
	}

	err := ctrl.deleteProduct.Execute(productId)

	if err != nil && errors.Is(err, application.ErrProductNotFound) {
		return utils.RespondWithError(c, err, 404)
	}

	if err != nil {
		return utils.RespondWithError(c, err, 500)
	}

	return utils.RespondWithSuccess(c, EmptyStruct{}, 200)
}
