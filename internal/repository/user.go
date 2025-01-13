package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/lautarok/yorcom/internal/domain"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repository *UserRepository) GetUserList() ([]domain.User, error) {
	var userList []domain.User
	err := repository.db.NewSelect().Model(&userList).Scan(context.Background())
	return userList, err
}

func (repository *UserRepository) GetByEmail(email string, user *domain.User) {
	err := repository.db.NewSelect().
		Model(user).
		Where("email = ?", email).
		Scan(context.Background())

	if err != sql.ErrNoRows && err != nil {
		log.Fatal(err)
	}
}

func (repository *UserRepository) CreateUser(user *domain.User) error {
	_, err := repository.db.NewInsert().Model(user).Exec(context.Background())
	return err
}
