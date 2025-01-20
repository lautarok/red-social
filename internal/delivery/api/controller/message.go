package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lautarok/yorcom/internal/service"
)

type MessageController struct {
	service   *service.MessageService
	wsService *service.WSService
}

func NewMessageController(service *service.MessageService, wsService *service.WSService) *MessageController {
	return &MessageController{
		service:   service,
		wsService: wsService,
	}
}

func (controller *MessageController) SendMessage(c *fiber.Ctx) error {
	body := struct {
		ConversationID int64  `json:"conversationId"`
		Message        string `json:"message"`
	}{}

	if err := c.BodyParser(&body); err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(map[string]string{
			"error": "wrong body",
		})
		return nil
	}

	fromId := c.Locals("auth_user_id").(int64)

	insertedId, userIds := controller.service.SendMessage(fromId, body.ConversationID, body.Message)

	message := controller.service.GetMessage(insertedId)

	for _, userId := range userIds {
		controller.wsService.Notify <- &service.WSNotify{
			ToUserID: userId,
			Data: map[string]interface{}{
				"type":    "message",
				"message": message,
			},
		}
	}

	c.Status(fiber.StatusCreated)
	c.JSON(map[string]int64{
		"insertedId": insertedId,
	})

	return nil
}
