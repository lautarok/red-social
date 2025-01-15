package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lautarok/yorcom/internal/service"
)

type MessageController struct {
	service *service.MessageService
}

func NewMessageController(service *service.MessageService) *MessageController {
	return &MessageController{service: service}
}

func (controller *MessageController) SendMessage(c *fiber.Ctx) error {
	body := struct {
		ConversationID int64  `json:"conversationId"`
		Message        string `json:"message"`
	}{}

	if err := c.BodyParser(&body); err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(map[string]string{
			"error": "wrong body",
		})
		return nil
	}

	fromId := c.Locals("auth_user_id").(int64)

	insertedId := controller.service.SendMessage(fromId, body.ConversationID, body.Message)

	c.Status(http.StatusCreated)
	c.JSON(map[string]int64{
		"insertedId": insertedId,
	})

	return nil
}
