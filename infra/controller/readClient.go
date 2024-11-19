package controller

import (
	"errors"
	"log"
	"luizalabs-challenge/application"
	"luizalabs-challenge/utils"

	"github.com/gofiber/fiber/v2"
)

type ReadClientController struct {
	readClient *application.ReadClient
}

func NewReadClientController(readClient *application.ReadClient) *ReadClientController {
	return &ReadClientController{
		readClient: readClient,
	}
}

func (ctrl *ReadClientController) Handle(c *fiber.Ctx) error {
	var input application.ReadClientInput

	err := c.ParamsParser(&input)

	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, fiber.ErrBadRequest, fiber.ErrBadRequest.Code)
	}

	if !utils.IsAuthorized(c, input.ClientId) {
		return utils.RespondWithError(c, fiber.ErrUnauthorized, fiber.ErrUnauthorized.Code)
	}

	client, err := ctrl.readClient.Execute(input)

	if err != nil && errors.Is(err, application.ErrClientNotFound) {
		log.Println(err)
		return utils.RespondWithError(c, err, 404)
	}

	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, fiber.ErrInternalServerError, 500)
	}

	return utils.RespondWithSuccess(c, client, 200)
}
