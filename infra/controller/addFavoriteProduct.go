package controller

import (
	"errors"
	"log"
	"luizalabs-challenge/application"
	"luizalabs-challenge/utils"

	"github.com/gofiber/fiber/v2"
)

type AddFavoriteProductController struct {
	addFavoriteProduct *application.AddFavoriteProduct
}

func NewAddFavoriteProductController(addFavoriteProduct *application.AddFavoriteProduct) *AddFavoriteProductController {
	return &AddFavoriteProductController{
		addFavoriteProduct: addFavoriteProduct,
	}
}

func (ctrl *AddFavoriteProductController) Handle(c *fiber.Ctx) error {
	clientId := c.Params("clientId")
	if clientId == "" {
		log.Println("ClientId não fornecido nos pathParams")
		return utils.RespondWithError(c, fiber.ErrBadRequest, fiber.ErrBadRequest.Code)
	}

	var favoriteProductInput application.AddFavoriteProductInput
	err := c.BodyParser(&favoriteProductInput)
	favoriteProductInput.ClientId = clientId

	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, fiber.ErrBadRequest, fiber.ErrBadRequest.Code)
	}

	if !utils.IsAuthorized(c, clientId) {
		return utils.RespondWithError(c, fiber.ErrUnauthorized, fiber.ErrUnauthorized.Code)
	}

	err = ctrl.addFavoriteProduct.Execute(favoriteProductInput)

	if err != nil && (errors.Is(err, application.ErrProductNotFound) || errors.Is(err, application.ErrClientNotFound)) {
		log.Println(err)
		return utils.RespondWithError(c, err, 404)
	}

	if err != nil && errors.Is(err, application.ErrProductIsAlreadyClientFavorite) {
		log.Println(err)
		return utils.RespondWithError(c, err, 400)
	}

	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, fiber.ErrInternalServerError, fiber.ErrInternalServerError.Code)
	}

	return utils.RespondWithSuccess(c, EmptyStruct{}, 201)
}
