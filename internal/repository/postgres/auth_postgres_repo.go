package postgres

import (
	"time"

	"github.com/evgeney-fullstack/auth-service/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user models.User) (int, error) {

	return 0, nil
}

func (r *AuthRepository) GetUser(email, username, password_hash string) (models.User, error) {

	return models.User{}, nil
}

func (r *AuthRepository) SaveRefreshTokenToDB(tokenID string, userID int, expiresAt time.Time) error {
	return nil
}

func (r *AuthRepository) IsRefreshTokenActive(tokenId string, userId int) (bool, error) {
	return true, nil
}

func (r *AuthRepository) DeleteRefreshTokenFromDB(tokenId string, userId int) error {
	return nil
}

func (r *AuthRepository) DeleteAllRefreshTokenFromDB(userId int) error {
	return nil
}
