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
	tx, err := repository.db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	result, err := tx.NewInsert().
		Model(message).
		Exec(context.Background())

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	var lastMessage domain.Message
	err = tx.NewSelect().
		Model(&lastMessage).
		Where("id = ?", insertedId).
		Scan(context.Background())
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = tx.NewUpdate().
		Model((*domain.Conversation)(nil)).
		Where("id = ?", message.ConversationID).
		Set("last_message = ?", lastMessage).
		Set("last_message_id = ?", lastMessage.ID).
		Exec(context.Background())
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return insertedId
}

func (repository *MessageRepository) GetMessage(id int64, message *domain.Message) error {
	return repository.db.NewSelect().
		Model(message).
		Where("id = ?", id).
		Scan(context.Background())
}
