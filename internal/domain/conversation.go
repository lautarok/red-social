package domain

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Conversation struct {
	ID        int64      `bun:",pk,autoincrement" json:"id"`
	Name      string     `json:"name"`
	Users     []User     `bun:"m2m:user_conversations,join:Conversation=User"`
	Messages  []*Message `bun:"rel:has-many,join:id=conversation_id"`
	CreatedAt time.Time  `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt time.Time  `bun:",nullzero,notnull,default:current_timestamp" json:"updatedAt"`
}

type UserConversation struct {
	UserID         int64         `bun:",pk"`
	User           *User         `bun:"rel:belongs-to,join:user_id=id"`
	ConversationID int64         `bun:",pk"`
	Conversation   *Conversation `bun:"rel:belongs-to,join:conversation_id=id"`
}

func (conversation *Conversation) BeforeInsert(ctx context.Context, _ bun.Query) error {
	conversation.CreatedAt = time.Now().UTC()
	conversation.UpdatedAt = time.Now().UTC()
	return nil
}

func (conversation *Conversation) BeforeUpdate(ctx context.Context, _ bun.Query) error {
	conversation.UpdatedAt = time.Now().UTC()
	return nil
}
