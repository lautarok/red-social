package service

import (
	"github.com/lautarok/yorcom/internal/domain"
	"github.com/lautarok/yorcom/internal/repository"
	"github.com/lautarok/yorcom/pkg/errors"
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

func (service *UserService) GetIDByEmail(email string) (int64, error) {
	var id int64
	ids, err := service.repository.GetIDSByEmail(email)
	if err != nil {
		return id, err
	} else if len(ids) == 0 {
		return id, errors.UserNotFound
	} else {
		id = ids[0]
	}
	return id, nil
}
