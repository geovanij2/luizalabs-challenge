package controller

import (
	"log"
	"luizalabs-challenge/application"
	"luizalabs-challenge/utils"

	"github.com/gofiber/fiber/v2"
)

type ListFavoriteProductsController struct {
	listFavoriteProducts *application.ListFavoriteProducts
}

func NewListFavoriteProductsController(listFavoriteProducts *application.ListFavoriteProducts) *ListFavoriteProductsController {
	return &ListFavoriteProductsController{
		listFavoriteProducts: listFavoriteProducts,
	}
}

func (ctrl *ListFavoriteProductsController) Handle(c *fiber.Ctx) error {
	clientId := c.Params("clientId")
	if clientId == "" {
		log.Println("ClientId não fornecido nos pathParams")
		return utils.RespondWithError(c, fiber.ErrBadRequest, fiber.ErrBadRequest.Code)
	}

	offset := c.QueryInt("offset", 0)
	if offset < 0 {
		log.Println("Offset inválido")
		return utils.RespondWithError(c, fiber.ErrBadRequest, fiber.ErrBadRequest.Code)
	}

	limit := c.QueryInt("limit", 10)
	if limit <= 0 {
		log.Println("Limit inválido")
		return utils.RespondWithError(c, fiber.ErrBadRequest, fiber.ErrBadRequest.Code)
	}

	listFavoriteProductsInput := application.ListFavoriteProductsInput{
		ClientId: clientId,
		Offset:   uint64(offset),
		Limit:    uint64(limit),
	}

	if !utils.IsAuthorized(c, listFavoriteProductsInput.ClientId) {
		return utils.RespondWithError(c, fiber.ErrUnauthorized, fiber.ErrUnauthorized.Code)
	}

	products, err := ctrl.listFavoriteProducts.Execute(listFavoriteProductsInput)
	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, fiber.ErrInternalServerError, fiber.ErrInternalServerError.Code)
	}

	return utils.RespondWithSuccess(c, products, 200)
}
