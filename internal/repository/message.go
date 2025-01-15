package repository

import (
	"context"
	"log"

	"github.com/lautarok/yorcom/internal/domain"
	"github.com/uptrace/bun"
)

type MessageRepository struct {
	db *bun.DB
}

func NewMessageRepository(db *bun.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (repository *MessageRepository) CreateMessage(message *domain.Message) int64 {
	result, err := repository.db.NewInsert().
		Model(message).
		Exec(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return insertedId
}
