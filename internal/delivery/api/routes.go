package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lautarok/yorcom/internal/delivery/api/controller"
	"github.com/lautarok/yorcom/internal/repository"
	"github.com/lautarok/yorcom/internal/service"
	"github.com/uptrace/bun"
)

func InitHttpRoutes(router fiber.Router, db *bun.DB) {
	router.Get("health/", func(c *fiber.Ctx) error {
		c.Status(http.StatusOK)
		c.JSON(map[string]string{
			"service": "backend",
			"status":  "alive",
		})
		return nil
	})

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	router.Get("user", userController.GetUserList)

	authService := service.NewAuthService(userRepository)
	authController := controller.NewAuthController(authService)
	router.Post("auth", authController.Login)
}
