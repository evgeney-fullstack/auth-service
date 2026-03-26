package postgres

import (
	"time"

	"github.com/evgeney-fullstack/auth-service/internal/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(email, username, password_hash string) (models.User, error)
	SaveRefreshTokenToDB(tokenID string, userID int, expiresAt time.Time) error
	IsRefreshTokenActive(tokenId string, userId int) (bool, error)
	DeleteRefreshTokenFromDB(tokenId string, userId int) error
	DeleteAllRefreshTokenFromDB(userId int) error
}

// Repository aggregates all store interfaces for database operations
type Repository struct {
	Authorization
}

// NewRepository constructs a new Repository with all available stores
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
