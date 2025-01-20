package domain

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Message struct {
	ID             int64     `bun:",pk,autoincrement" json:"id"`
	Text           string    `json:"text"`
	ConversationID int64     `json:"conversationId"`
	FromUserID     int64     `json:"fromUserId"`
	CreatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updatedAt"`
}

func (message *Message) BeforeInsert(ctx context.Context, _ bun.Query) error {
	message.CreatedAt = time.Now().UTC()
	message.UpdatedAt = time.Now().UTC()
	return nil
}

func (message *Message) BeforeUpdate(ctx context.Context, _ bun.Query) error {
	message.UpdatedAt = time.Now().UTC()
	return nil
}
