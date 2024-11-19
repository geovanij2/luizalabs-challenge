package controller

import (
	"errors"
	"log"
	"luizalabs-chalenge/application"
	"luizalabs-chalenge/utils"

	"github.com/gofiber/fiber/v2"
)

type LoginController struct {
	login *application.Login
}

type LoginReponse struct {
	AccessToken string `json:"accessToken"`
}

func NewLoginController(login *application.Login) *LoginController {
	return &LoginController{
		login: login,
	}
}

func (ctrl *LoginController) Handle(c *fiber.Ctx) error {
	var loginInput application.LoginInput
	err := c.BodyParser(&loginInput)

	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, fiber.ErrBadRequest, fiber.ErrBadRequest.Code)
	}

	accessToken, err := ctrl.login.Execute(loginInput)

	if err != nil && errors.Is(err, application.ErrClientNotFound) {
		log.Println(err)
		return utils.RespondWithError(c, err, 404)
	}

	if err != nil && errors.Is(err, application.ErrWrongPassword) {
		log.Println(err)
		return utils.RespondWithError(c, err, 400)
	}

	if err != nil {
		log.Println(err)
		return utils.RespondWithError(c, fiber.ErrInternalServerError, fiber.ErrInternalServerError.Code)
	}

	return utils.RespondWithSuccess(c, LoginReponse{
		AccessToken: accessToken,
	}, 201)
}
