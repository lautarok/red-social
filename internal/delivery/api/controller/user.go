package controller

import (
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

func (controller *UserController) GetUserList(c *fiber.Ctx) error {
	user_list, err := controller.service.GetUserList()
	if err != nil {
		c.Status(500)
		c.JSON(map[string]string{
			"error": "internal error",
		})
		return err
	}

	c.Status(200)
	c.JSON(map[string]interface{}{
		"userList": user_list,
	})

	return nil
}
