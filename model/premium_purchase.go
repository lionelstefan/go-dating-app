package model

import (
	"time"

	"github.com/google/uuid"
)

type PremiumPurchase struct {
	ID uuid.UUID `db:"id"`
	UserId uuid.UUID `db:"user_id"`
	PremiumType string `db:"premium_type"`
	CreatedAt time.Time `db:"created_at"` 
	UpdatedAt time.Time `db:"updated_at"` 
}