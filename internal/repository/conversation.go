package repository

import (
	"context"
	"log"

	"github.com/lautarok/yorcom/internal/domain"
	"github.com/uptrace/bun"
)

type ConversationRepository struct {
	db *bun.DB
}

func NewConversationRepository(db *bun.DB) *ConversationRepository {
	return &ConversationRepository{
		db: db,
	}
}

func (repository *ConversationRepository) GetIDsByEmails(emails ...string) ([]int64, error) {
	var ids []int64
	var users []domain.User

	err := repository.db.NewSelect().
		Column("id").
		Model(&users).
		Where("email in (?)", bun.In(emails)).
		Scan(context.Background())

	if err != nil {
		return ids, err
	}

	for _, u := range users {
		ids = append(ids, u.ID)
	}

	return ids, nil
}

func (repository *ConversationRepository) ConversationExists(userIds ...int64) (bool, error) {
	return repository.db.NewSelect().
		Model((*domain.Conversation)(nil)).
		Join("JOIN user_conversations uc ON uc.conversation_id = conversation.id").
		Where("uc.user_id in (?)", bun.In(userIds)).
		Group("conversation.id").
		Having("COUNT(DISTINCT uc.user_id) = ?", len(userIds)).
		Exists(context.Background())
}

func (repository *ConversationRepository) CreateConversation(conversation *domain.Conversation) (int64, error) {
	result, err := repository.db.NewInsert().
		Model(conversation).
		Exec(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return lastInsertId, nil
}

func (repository *ConversationRepository) AppendUsersToConversation(conversationId int64, userIds ...int64) error {
	tx, err := repository.db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, userId := range userIds {
		_, err := tx.NewInsert().
			Model(&domain.UserConversation{
				UserID:         userId,
				ConversationID: conversationId,
			}).
			Exec(context.Background())

		if err != nil {
			tx.Rollback()
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
