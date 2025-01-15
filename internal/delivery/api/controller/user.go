package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lautarok/yorcom/internal/service"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (controller *UserController) GetMyUser(c *fiber.Ctx) error {
	myUser := controller.service.GetUserByID(c.Locals("auth_user_id").(int64))

	c.Status(http.StatusOK)
	c.JSON(myUser)

	return nil
}
