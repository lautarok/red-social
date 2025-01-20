package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/lautarok/yorcom/internal/domain"
	"github.com/lautarok/yorcom/pkg/errors"
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
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (repository *ConversationRepository) GetConversationList(userId int64, conversationList *[]domain.Conversation) error {
	return repository.db.NewSelect().
		Model(conversationList).
		Join("JOIN user_conversations ON user_conversations.conversation_id = conversation.id").
		Where("user_conversations.user_id = ?", userId).
		OrderExpr("CASE WHEN last_message_id > 0 THEN last_message_id ELSE id END DESC").
		Relation("Users").
		Scan(context.Background())
}

func (repository *ConversationRepository) GetByID(id int64, conversation *domain.Conversation) error {
	err := repository.db.NewSelect().
		Model(conversation).
		Where("id = ?", id).
		Relation("Messages", func(sq *bun.SelectQuery) *bun.SelectQuery {
			return sq.Order("id ASC").Limit(25)
		}).
		Relation("Users").
		Scan(context.Background())

	if err == sql.ErrNoRows {
		return errors.ConversationNotFound
	}

	return err
}

func (repository *ConversationRepository) GetUsers(id int64, users *[]domain.User) error {
	var conversation domain.Conversation

	err := repository.db.NewSelect().
		Model(&conversation).
		Where("id = ?", id).
		Relation("Users").
		Scan(context.Background())

	if err != nil {
		return err
	}

	*users = conversation.Users

	return nil
}
