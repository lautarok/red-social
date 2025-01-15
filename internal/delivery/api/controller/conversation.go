package controller

import (
	"log"
	"net/http"
	"strconv"

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
		ID string `json:"id"`
	}{}

	if err := c.BodyParser(&body); err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(map[string]string{
			"error": "wrong body",
		})
		return nil
	} else if !util.ValidateID(body.ID) {
		c.Status(http.StatusBadRequest)
		c.JSON(map[string]string{
			"error": "wrong id",
		})
		return nil
	}

	id, _ := strconv.Atoi(body.ID)

	conversationId, err := controller.service.CreateConversation(c.Locals("auth_user_id").(int64), int64(id))
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

func (controller *ConversationController) GetConversationList(c *fiber.Ctx) error {
	authUserId, ok := c.Locals("auth_user_id").(int64)
	if !ok {
		c.SendStatus(http.StatusUnauthorized)
		return nil
	}

	conversationList, err := controller.service.GetConversationList(authUserId)
	if err != nil {
		c.SendStatus(http.StatusInternalServerError)
		log.Fatal(err)
		return nil
	}

	c.Status(http.StatusOK)
	c.JSON(conversationList)

	return nil
}

func (controller *ConversationController) GetConversation(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(map[string]string{
			"error": "wrong id param",
		})
		return nil
	}

	conversation, err := controller.service.GetConversation(int64(id))
	if err != nil {
		if err == errors.ConversationNotFound {
			c.Status(http.StatusNotFound)
			c.JSON(map[string]string{
				"error": "conversation not found",
			})
			return nil
		} else {
			log.Fatal(err)
			return nil
		}
	}

	c.Status(http.StatusOK)
	c.JSON(conversation)

	return nil
}
