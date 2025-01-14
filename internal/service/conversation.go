package service

import (
	"github.com/lautarok/yorcom/internal/domain"
	"github.com/lautarok/yorcom/internal/repository"
	"github.com/lautarok/yorcom/pkg/errors"
)

type ConversationService struct {
	repository *repository.ConversationRepository
}

func NewConversationService(repository *repository.ConversationRepository) *ConversationService {
	return &ConversationService{
		repository: repository,
	}
}

func (service *ConversationService) CreateConversation(userAEmail string, userBEmail string) (int64, error) {
	var conversationId int64

	userIds, err := service.repository.GetIDsByEmails(userAEmail, userBEmail)
	if err != nil {
		return conversationId, err
	}

	if len(userIds) < 2 {
		return conversationId, errors.UserNotFound
	}

	exists, err := service.repository.ConversationExists(userIds...)
	if err != nil {
		return conversationId, err
	} else if exists {
		return conversationId, errors.ConversationAlreadyExists
	}

	conversationId, err = service.repository.CreateConversation(&domain.Conversation{
		Users: []domain.User{
			{
				ID: userIds[0],
			},
			{
				ID: userIds[1],
			},
		},
	})

	service.repository.AppendUsersToConversation(conversationId, userIds...)

	return conversationId, err
}
