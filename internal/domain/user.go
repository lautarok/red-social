package domain

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	ID         int64     `bun:",pk,autoincrement" json:"id"`
	GivenName  string    `json:"givenName"`
	FamilyName string    `json:"familyName"`
	Password   string    `json:"-"`
	Email      string    `bun:",unique" json:"email"`
	CreatedAt  time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt  time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updatedAt"`
}

func (user *User) BeforeInsert(ctx context.Context, _ bun.Query) error {
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()
	return nil
}

func (user *User) BeforeUpdate(ctx context.Context, _ bun.Query) error {
	user.UpdatedAt = time.Now().UTC()
	return nil
}
