package model

import (
	"time"

	"github.com/google/uuid"
)

type UserRegister struct {
	Email string `db:"email"`
	Password string `db:"password"`
	FullName string `db:"full_name"`
	Age string `db:"age"`
	Gender string `db:"gender"`
	IsVerified bool `db:"is_verified"`
	NoSwipeLimit bool `db:"no_swipe_limit"`
}

type User struct {
	Id uuid.UUID `db:"id"`
	Email string `db:"email"`
	Password string `db:"password"`
	FullName string `db:"full_name"`
	Age string `db:"age"`
	Gender string `db:"gender"`
	IsVerified bool `db:"is_premium"`
	NoSwipeLimit bool `db:"no_swipe_limit"`
	CreatedAt time.Time `db:"created_at"` 
	UpdatedAt time.Time `db:"updated_at"` 
}