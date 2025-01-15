package service

import (
	"github.com/lautarok/yorcom/internal/domain"
	"github.com/lautarok/yorcom/internal/repository"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (service *UserService) GetUserByID(id int64) domain.User {
	var user domain.User
	service.repository.GetByID(id, &user)
	return user
}
