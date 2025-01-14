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

type requestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

func (controller *AuthController) SignUp(c *fiber.Ctx) error {
	body := struct {
		GivenName  string `json:"givenName"`
		FamilyName string `json:"familyName"`
		Email      string `json:"email"`
		Password   string `json:"password"`
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
	} else if len(body.GivenName) < 3 || len(body.GivenName) > 24 {
		c.Status(http.StatusBadRequest)
		c.JSON(map[string]string{
			"error": "wrong given name",
		})
		return nil
	} else if len(body.FamilyName) < 3 || len(body.FamilyName) > 24 {
		c.Status(http.StatusBadRequest)
		c.JSON(map[string]string{
			"error": "wrong family name",
		})
		return nil
	} else if !util.ValidateEmail(body.Email) {
		c.Status(http.StatusBadRequest)
		c.JSON(map[string]string{
			"error": "wrong email",
		})
		return nil
	}

	token, err := controller.service.SignUp(
		body.Email,
		body.Password,
		body.GivenName,
		body.FamilyName,
	)
	if err == errors.UserAlreadyExists {
		c.Status(http.StatusConflict)
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
