package controller

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/lautarok/yorcom/internal/service"
)

type WSController struct {
	service *service.WSService
}

func NewWSController(service *service.WSService) *WSController {
	return &WSController{service: service}
}

func (controller *WSController) Connect(c *websocket.Conn) {
	controller.service.Subscribe <- &service.WSClient{
		Context: c,
		UserID:  c.Locals("auth_user_id").(int64),
	}
	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			break
		}
	}
}
