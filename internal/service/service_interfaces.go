package service

import (
	"github.com/evgeney-fullstack/auth-service/internal/models"
	"github.com/evgeney-fullstack/auth-service/internal/repository/postgres"
	"github.com/evgeney-fullstack/auth-service/internal/repository/redis"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateAccessToken(email, username, password_hash string) (string, error)
	UpdateAccessToken(userId int) (string, error)
	RevokingAllAccessToken(userId int) error
	GenerateRefreshToken(email, username, password_hash string) (string, error)
	ParseAccessToken(accessToken string) (int, error)
	ParseRefreshToken(refreshToken string) (string, int, error)
	IsRefreshTokenActive(tokenId string, userId int) (bool, error)
	DeleteRefreshToken(tokenId string, userId int) error
	DeleteAllRefreshToken(userId int) error
}

// Service layer aggregates all business logic services
type Service struct {
	Authorization
}

// NewService constructs new Service layer with business logic
func NewService(repos *postgres.Repository, cacheRepo *redis.CacheRepository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, cacheRepo.Authorization),
	}
}
