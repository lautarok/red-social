package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (controller *HealthController) GetHealth(c *fiber.Ctx) error {
	c.Status(http.StatusOK)
	c.JSON(map[string]string{
		"service": "backend",
		"status":  "alive",
	})
	return nil
}
