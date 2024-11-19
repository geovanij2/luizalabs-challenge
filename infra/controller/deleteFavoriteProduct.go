package controller

import (
	"errors"
	"log"
	"luizalabs-chalenge/application"
	"luizalabs-chalenge/utils"

	"github.com/gofiber/fiber/v2"
)

type DeleteFavoriteProductController struct {
	deleteFavoriteProduct *application.DeleteFavoriteProduct
}

func NewDeleteFavoriteProductController(deleteFavoriteProduct *application.DeleteFavoriteProduct) *DeleteFavoriteProductController {
	return &DeleteFavoriteProductController{
		deleteFavoriteProduct: deleteFavoriteProduct,
	}
}

func (ctrl *DeleteFavoriteProductController) Handle(c *fiber.Ctx) error {
	var deleteFavoriteProductInput application.DeleteFavoriteProductInput
	err := c.ParamsParser(&deleteFavoriteProductInput)

	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, fiber.ErrBadRequest, fiber.ErrBadRequest.Code)
	}

	if !utils.IsAuthorized(c, deleteFavoriteProductInput.ClientId) {
		return utils.RespondWithError(c, fiber.ErrUnauthorized, fiber.ErrUnauthorized.Code)
	}

	err = ctrl.deleteFavoriteProduct.Execute(deleteFavoriteProductInput)

	if err != nil && (errors.Is(err, application.ErrProductNotFound) || errors.Is(err, application.ErrClientNotFound)) {
		log.Println(err)
		return utils.RespondWithError(c, err, 404)
	}

	if err != nil && errors.Is(err, application.ErrProductIsNotFavorite) {
		log.Println(err)
		return utils.RespondWithError(c, err, 400)
	}

	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, fiber.ErrInternalServerError, 500)
	}

	return utils.RespondWithSuccess(c, EmptyStruct{}, 200)
}
