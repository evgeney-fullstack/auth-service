package models

import (
	"time"
)

type User struct {
	Id           int       `json:"-" db:"id" bson:"_id"`
	Email        string    `json:"email" bson:"email" binding:"required"`
	PasswordHash string    `json:"password_hash" bson:"password_hash" binding:"required"`
	Username     string    `json:"username" bson:"username" binding:"required"`
	IsVerified   bool      `json:"is_verified" bson:"is_verified"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}

type RefreshToken struct {
	TokenId   string    `json:"token_id" bson:"token_id"`
	UserId    int       `json:"user_id" bson:"user_id"`
	ExpiresAt time.Time `json:"expires_at" bson:"expires_at" db:"expires_at"`
}
