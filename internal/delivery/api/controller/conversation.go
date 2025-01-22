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
	service     *service.ConversationService
	wsService   *service.WSService
	userService *service.UserService
}

func NewConversationController(service *service.ConversationService, wsService *service.WSService, userService *service.UserService) *ConversationController {
	return &ConversationController{
		service:     service,
		wsService:   wsService,
		userService: userService,
	}
}

func (controller *ConversationController) CreateConversation(c *fiber.Ctx) error {
	body := struct {
		Email string `json:"email"`
	}{}

	if err := c.BodyParser(&body); err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(map[string]string{
			"error": "wrong body",
		})
		return nil
	} else if !util.ValidateEmail(body.Email) {
		c.Status(fiber.StatusBadRequest)
		c.JSON(map[string]string{
			"error": "wrong email",
		})
		return nil
	}

	myUserEmail := c.Locals("auth_user_email").(string)
	myUserId := c.Locals("auth_user_id").(int64)

	conversationId, err := controller.service.CreateConversation(myUserEmail, body.Email)
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

	toUsersId, err := controller.userService.GetIDByEmail(body.Email)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	controller.wsService.Notify <- &service.WSNotify{
		ToUserID: int64(toUsersId),
		Data: map[string]interface{}{
			"type":         "conversation",
			"conversation": conversation,
		},
	}

	controller.wsService.Notify <- &service.WSNotify{
		ToUserID: myUserId,
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
