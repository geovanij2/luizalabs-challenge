package controller

import (
	"errors"
	"log"
	"luizalabs-challenge/application"
	"luizalabs-challenge/domain/entity"
	"luizalabs-challenge/utils"

	"github.com/gofiber/fiber/v2"
)

type UpdateClientController struct {
	updateClient *application.UpdateClient
}

func NewUpdateClientController(updateClient *application.UpdateClient) *UpdateClientController {
	return &UpdateClientController{
		updateClient: updateClient,
	}
}

func (ctrl *UpdateClientController) Handle(c *fiber.Ctx) error {
	clientId := c.Params("clientId")
	if clientId == "" {
		log.Println("ClientId não fornecido nos pathParams")
		return utils.RespondWithError(c, fiber.ErrBadRequest, fiber.ErrBadRequest.Code)
	}

	var input entity.Client

	err := c.BodyParser(&input)
	input.Id = clientId

	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, fiber.ErrBadRequest, fiber.ErrBadRequest.Code)
	}

	if !utils.IsAuthorized(c, input.Id) {
		return utils.RespondWithError(c, fiber.ErrUnauthorized, fiber.ErrUnauthorized.Code)
	}

	err = ctrl.updateClient.Execute(&input)

	if err != nil && errors.Is(err, application.ErrClientNotFound) {
		log.Println(err)
		return utils.RespondWithError(c, err, 404)
	}

	if err != nil && errors.Is(err, application.ErrClientEmailAlreadyExists) {
		log.Println(err)
		return utils.RespondWithError(c, err, 400)
	}

	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, err, 500)
	}

	return utils.RespondWithSuccess(c, input, 200)
}
