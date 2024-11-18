package controller

import (
	"errors"
	"log"
	"luizalabs-chalenge/application"
	"luizalabs-chalenge/domain/entity"
	"luizalabs-chalenge/utils"

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
	var input entity.Client

	err := c.BodyParser(&input)

	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, fiber.ErrBadRequest, fiber.ErrBadRequest.Code)
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
