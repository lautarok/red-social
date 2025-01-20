package service

import (
	"log"

	"github.com/lautarok/yorcom/internal/domain"
	"github.com/lautarok/yorcom/internal/repository"
)

type MessageService struct {
	repository             *repository.MessageRepository
	conversationRepository *repository.ConversationRepository
}

func NewMessageService(repository *repository.MessageRepository, conversationRepository *repository.ConversationRepository) *MessageService {
	return &MessageService{
		repository:             repository,
		conversationRepository: conversationRepository,
	}
}

func (service *MessageService) SendMessage(fromId int64, toConversationId int64, message string) (int64, []int64) {
	insertedId := service.repository.CreateMessage(&domain.Message{
		FromUserID:     fromId,
		ConversationID: toConversationId,
		Text:           message,
	})

	var users []domain.User
	err := service.conversationRepository.GetUsers(toConversationId, &users)
	if err != nil {
		log.Fatal(err)
	}

	var userIds []int64
	for _, user := range users {
		userIds = append(userIds, user.ID)
	}

	return insertedId, userIds
}

func (service *MessageService) GetMessage(id int64) domain.Message {
	var message domain.Message
	err := service.repository.GetMessage(id, &message)
	if err != nil {
		log.Fatal(err)
	}
	return message
}
