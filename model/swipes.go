package model

import (
	"time"

	"github.com/google/uuid"
)

type Swipes struct {
	ID uuid.UUID `db:"id"`
	UserId uuid.UUID `db:"user_id"`
	TargetId uuid.UUID `db:"target_Id"`
	IsLike bool `db:"is_like"`
	CreatedAt time.Time `db:"created_at"` 
	UpdatedAt time.Time `db:"updated_at"` 
}