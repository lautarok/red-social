package controller

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lautarok/yorcom/internal/service"
	"github.com/lautarok/yorcom/pkg/errors"
	"github.com/lautarok/yorcom/pkg/util"
)

type ConversationController struct {
	service *service.ConversationService
}

func NewConversationController(service *service.ConversationService) *ConversationController {
	return &ConversationController{
		service: service,
	}
}

func (controller *ConversationController) CreateConversation(c *fiber.Ctx) error {
	body := struct {
		Email string `json:"email"`
	}{}

	if err := c.BodyParser(&body); err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(map[string]string{
			"error": "wrong body",
		})
		return nil
	} else if !util.ValidateEmail(body.Email) {
		c.Status(http.StatusBadRequest)
		c.JSON(map[string]string{
			"error": "wrong email",
		})
		return nil
	}

	conversationId, err := controller.service.CreateConversation(c.Locals("auth_user_email").(string), body.Email)
	if err != nil {
		if err == errors.UserNotFound {
			c.Status(http.StatusNotFound)
			c.JSON(map[string]string{
				"error": err.Error(),
			})
			return nil
		} else if err == errors.ConversationAlreadyExists {
			c.Status(http.StatusConflict)
			c.JSON(map[string]string{
				"error": err.Error(),
			})
			return nil
		} else {
			log.Fatal(err)
			return nil
		}
	}

	c.Status(http.StatusCreated)
	c.JSON(map[string]int64{
		"conversationId": conversationId,
	})

	return nil
}
