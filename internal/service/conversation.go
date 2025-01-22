package service

import (
	"github.com/lautarok/yorcom/internal/domain"
	"github.com/lautarok/yorcom/internal/repository"
	"github.com/lautarok/yorcom/pkg/errors"
)

type ConversationService struct {
	repository     *repository.ConversationRepository
	userRepository *repository.UserRepository
}

func NewConversationService(repository *repository.ConversationRepository, userRepository *repository.UserRepository) *ConversationService {
	return &ConversationService{
		repository:     repository,
		userRepository: userRepository,
	}
}

func (service *ConversationService) CreateConversation(userEmails ...string) (int64, error) {
	var conversationId int64

	userIds, err := service.userRepository.GetIDSByEmail(userEmails...)
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
	if err != nil {
		return conversationId, err
	}

	service.repository.AppendUsersToConversation(conversationId, userIds...)

	return conversationId, err
}

func (service *ConversationService) GetConversationList(userId int64) ([]domain.Conversation, error) {
	var conversationList []domain.Conversation

	err := service.repository.GetConversationList(userId, &conversationList)
	if err != nil {
		return conversationList, err
	}

	return conversationList, nil
}

func (service *ConversationService) GetConversation(id int64) (domain.Conversation, error) {
	var conversation domain.Conversation
	err := service.repository.GetByID(id, &conversation)
	return conversation, err
}
