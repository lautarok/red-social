package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/lautarok/yorcom/internal/domain"
	"github.com/lautarok/yorcom/pkg/errors"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repository *UserRepository) GetByID(id int64, user *domain.User) {
	err := repository.db.NewSelect().
		Model(user).
		Where("id = ?", id).
		Scan(context.Background())

	if err != sql.ErrNoRows && err != nil {
		log.Fatal(err)
	}
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

func (repository *UserRepository) GetIDSByEmail(emails ...string) ([]int64, error) {
	var ids []int64

	var users []domain.User
	err := repository.db.NewSelect().
		Model(&users).
		Where("email in (?)", bun.In(emails)).
		Scan(context.Background())

	if err != nil {
		return ids, err
	}

	for _, user := range users {
		ids = append(ids, user.ID)
	}

	return ids, nil
}

func (repository *UserRepository) CreateUser(user *domain.User) (int64, error) {
	var insertedId int64

	result, err := repository.db.NewInsert().Model(user).Exec(context.Background())

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return insertedId, errors.UserAlreadyExists
		}
		return insertedId, err
	}

	insertedId, err = result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return insertedId, nil
}
