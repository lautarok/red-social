package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lautarok/yorcom/internal/delivery/api/controller"
	"github.com/lautarok/yorcom/internal/delivery/api/middleware"
	"github.com/lautarok/yorcom/internal/repository"
	"github.com/lautarok/yorcom/internal/service"
	"github.com/uptrace/bun"
)

func InitHttpRoutes(router fiber.Router, db *bun.DB) {
	healthController := controller.NewHealthController()
	router.Get("health/", healthController.GetHealth)

	authMiddleware := middleware.NewAuthMiddleware()
	userRepository := repository.NewUserRepository(db)
	conversationRepository := repository.NewConversationRepository(db)
	conversationService := service.NewConversationService(conversationRepository)
	authService := service.NewAuthService(userRepository)
	conversationController := controller.NewConversationController(conversationService)

	authController := controller.NewAuthController(authService)
	router.Post("auth", authController.Login)
	router.Put("auth/sign-up", authController.SignUp)

	router.Post("conversation", authMiddleware.IsUser(conversationController.CreateConversation))
}
