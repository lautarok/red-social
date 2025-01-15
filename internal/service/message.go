package service

import (
	"github.com/lautarok/yorcom/internal/domain"
	"github.com/lautarok/yorcom/internal/repository"
)

type MessageService struct {
	repository *repository.MessageRepository
}

func NewMessageService(repository *repository.MessageRepository) *MessageService {
	return &MessageService{repository: repository}
}

func (service *MessageService) SendMessage(fromId int64, toConversationId int64, message string) int64 {
	return service.repository.CreateMessage(&domain.Message{
		FromUserID:     fromId,
		ConversationID: toConversationId,
		Text:           message,
	})
}
