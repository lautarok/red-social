package controller

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lautarok/yorcom/internal/service"
	"github.com/lautarok/yorcom/pkg/errors"
	"github.com/lautarok/yorcom/pkg/util"
)

type ConversationController struct {
	service   *service.ConversationService
	wsService *service.WSService
}

func NewConversationController(service *service.ConversationService, wsService *service.WSService) *ConversationController {
	return &ConversationController{
		service:   service,
		wsService: wsService,
	}
}

func (controller *ConversationController) CreateConversation(c *fiber.Ctx) error {
	body := struct {
		ID string `json:"id"`
	}{}

	if err := c.BodyParser(&body); err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(map[string]string{
			"error": "wrong body",
		})
		return nil
	} else if !util.ValidateID(body.ID) {
		c.Status(fiber.StatusBadRequest)
		c.JSON(map[string]string{
			"error": "wrong id",
		})
		return nil
	}

	id, _ := strconv.Atoi(body.ID)

	myUserId := c.Locals("auth_user_id").(int64)

	conversationId, err := controller.service.CreateConversation(myUserId, int64(id))
	if err != nil {
		if err == errors.UserNotFound {
			c.Status(fiber.StatusNotFound)
			c.JSON(map[string]string{
				"error": err.Error(),
			})
			return nil
		} else if err == errors.ConversationAlreadyExists {
			c.Status(fiber.StatusConflict)
			c.JSON(map[string]string{
				"error": err.Error(),
			})
			return nil
		} else {
			log.Fatal(err)
			return nil
		}
	}

	conversation, err := controller.service.GetConversation(conversationId)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	controller.wsService.Notify <- &service.WSNotify{
		ToUserID: int64(id),
		Data: map[string]interface{}{
			"type":         "conversation",
			"conversation": conversation,
		},
	}

	c.Status(fiber.StatusCreated)
	c.JSON(conversation)

	return nil
}

func (controller *ConversationController) GetConversationList(c *fiber.Ctx) error {
	authUserId, ok := c.Locals("auth_user_id").(int64)
	if !ok {
		c.SendStatus(fiber.StatusUnauthorized)
		return nil
	}

	conversationList, err := controller.service.GetConversationList(authUserId)
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		log.Fatal(err)
		return nil
	}

	c.Status(fiber.StatusOK)
	c.JSON(conversationList)

	return nil
}

func (controller *ConversationController) GetConversation(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(map[string]string{
			"error": "wrong id param",
		})
		return nil
	}

	conversation, err := controller.service.GetConversation(int64(id))
	if err != nil {
		if err == errors.ConversationNotFound {
			c.Status(fiber.StatusNotFound)
			c.JSON(map[string]string{
				"error": "conversation not found",
			})
			return nil
		} else {
			log.Fatal(err)
			return nil
		}
	}

	c.Status(fiber.StatusOK)
	c.JSON(conversation)

	return nil
}
