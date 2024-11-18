package controller

import (
	"errors"
	"log"
	"luizalabs-chalenge/application"
	"luizalabs-chalenge/utils"

	"github.com/gofiber/fiber/v2"
)

type CreateClientController struct {
	createClient *application.CreateClient
}

func NewCreateClientController(createClient *application.CreateClient) *CreateClientController {
	return &CreateClientController{
		createClient: createClient,
	}
}

func (ctrl *CreateClientController) Handle(c *fiber.Ctx) error {
	var clientInput application.CreateClientInput
	err := c.BodyParser(&clientInput)

	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, fiber.ErrBadRequest, fiber.ErrBadRequest.Code)
	}

	client, err := ctrl.createClient.Execute(clientInput)

	if err != nil && errors.Is(err, application.ErrClientEmailAlreadyExists) {
		log.Println(err)
		return utils.RespondWithError(c, err, 400)
	}

	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, fiber.ErrInternalServerError, fiber.ErrInternalServerError.Code)
	}

	return utils.RespondWithSuccess(c, client, 201)
}
