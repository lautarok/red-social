package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lautarok/yorcom/internal/service"
	"github.com/lautarok/yorcom/pkg/errors"
	"github.com/lautarok/yorcom/pkg/util"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (controller *AuthController) Login(c *fiber.Ctx) error {
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&body); err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(map[string]string{
			"error": "wrong body",
		})
		return nil
	} else if len(body.Password) < 6 || len(body.Password) > 16 {
		c.Status(http.StatusBadRequest)
		c.JSON(map[string]string{
			"error": "wrong password",
		})
		return nil
	} else if !util.ValidateEmail(body.Email) {
		c.Status(http.StatusBadRequest)
		c.JSON(map[string]string{
			"error": "wrong email",
		})
		return nil
	}

	token, err := controller.service.Login(body.Email, body.Password)
	if err == errors.UserNotFound {
		c.Status(http.StatusNotFound)
		c.JSON(map[string]string{
			"error": err.Error(),
		})
		return nil
	} else if err == errors.InvalidPassword {
		c.Status(http.StatusUnauthorized)
		c.JSON(map[string]string{
			"error": err.Error(),
		})
		return nil
	} else if err != nil {
		return err
	}

	c.Status(http.StatusCreated)
	c.JSON(map[string]string{
		"token": token,
	})

	return nil
}
