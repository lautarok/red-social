package service

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lautarok/yorcom/internal/domain"
	"github.com/lautarok/yorcom/internal/repository"
	"github.com/lautarok/yorcom/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (service *AuthService) Login(email string, password string) (string, error) {
	var user domain.User
	var token string

	service.userRepository.GetByEmail(email, &user)

	if user.ID == 0 {
		return token, errors.UserNotFound
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return token, errors.InvalidPassword
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
	})

	token, err = jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return token, err
	}

	return token, nil
}
