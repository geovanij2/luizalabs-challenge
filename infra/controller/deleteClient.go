package controller

import (
	"errors"
	"log"
	"luizalabs-chalenge/application"
	"luizalabs-chalenge/utils"

	"github.com/gofiber/fiber/v2"
)

type DeleteClientController struct {
	deleteClient *application.DeleteClient
}

type EmptyStruct struct{}

func NewDeleteClientController(deleteClient *application.DeleteClient) *DeleteClientController {
	return &DeleteClientController{
		deleteClient: deleteClient,
	}
}

func (ctrl *DeleteClientController) Handle(c *fiber.Ctx) error {
	clientId := c.Params("clientId")

	if clientId == "" {
		return utils.RespondWithError(c, ErrMissingParameter, 400)
	}

	err := ctrl.deleteClient.Execute(clientId)

	if err != nil && errors.Is(err, application.ErrClientNotFound) {
		log.Println(err)
		return utils.RespondWithError(c, err, 404)
	}

	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, fiber.ErrInternalServerError, 500)
	}

	return utils.RespondWithSuccess(c, EmptyStruct{}, 200)
}
