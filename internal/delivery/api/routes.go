package api

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/lautarok/yorcom/internal/delivery/api/controller"
	"github.com/lautarok/yorcom/internal/delivery/api/middleware"
	"github.com/lautarok/yorcom/internal/repository"
	"github.com/lautarok/yorcom/internal/service"
	"github.com/uptrace/bun"
)

func InitRoutes(router fiber.Router, db *bun.DB) {
	authMiddleware := middleware.NewAuthMiddleware()
	userRepository := repository.NewUserRepository(db)
	conversationRepository := repository.NewConversationRepository(db)
	messageRepository := repository.NewMessageRepository(db)
	conversationService := service.NewConversationService(conversationRepository)
	authService := service.NewAuthService(userRepository)
	userService := service.NewUserService(userRepository)
	messageService := service.NewMessageService(messageRepository, conversationRepository)
	wsService := service.NewWSService()
	conversationController := controller.NewConversationController(conversationService, wsService)
	authController := controller.NewAuthController(authService)
	userController := controller.NewUserController(userService)
	messageController := controller.NewMessageController(messageService, wsService)
	healthController := controller.NewHealthController()
	wsController := controller.NewWSController(wsService)

	router.Get("health/", healthController.GetHealth)

	router.Post("auth", authController.Login)
	router.Put("auth/sign-up", authController.SignUp)

	router.Get("user/me", authMiddleware.IsUser(userController.GetMyUser))

	router.Post("conversation", authMiddleware.IsUser(conversationController.CreateConversation))
	router.Get("conversation", authMiddleware.IsUser(conversationController.GetConversationList))
	router.Get("conversation/:id", authMiddleware.IsUser(conversationController.GetConversation))

	router.Post("message", authMiddleware.IsUser(messageController.SendMessage))

	router.Get("ws", authMiddleware.IsUser(websocket.New(wsController.Connect)))
}
