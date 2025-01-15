package service

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lautarok/yorcom/internal/domain"
	"github.com/lautarok/yorcom/internal/repository"
	"github.com/lautarok/yorcom/pkg/errors"
	"github.com/lautarok/yorcom/pkg/util"
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

func (service *AuthService) generateToken(email string, id int64) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
	})

	token, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return token, err
	}

	return token, nil
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

	token, err = service.generateToken(user.Email, user.ID)

	return token, nil
}

func (service *AuthService) SignUp(email string, password string, givenName string, familyName string) (string, error) {
	var token string

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	id, err := service.userRepository.CreateUser(&domain.User{
		Email:      email,
		Password:   string(passwordHash),
		FamilyName: util.Capitalize(familyName),
		GivenName:  util.Capitalize(givenName),
	})

	if err == errors.UserAlreadyExists {
		return token, err
	} else if err != nil {
		log.Fatal(err)
	}

	token, err = service.generateToken(email, id)
	if err != nil {
		log.Fatal(err)
	}

	return token, nil
}
