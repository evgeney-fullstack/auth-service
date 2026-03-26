package service

import (
	"time"

	"github.com/evgeney-fullstack/auth-service/internal/models"
	"github.com/evgeney-fullstack/auth-service/internal/repository/postgres"
	"github.com/evgeney-fullstack/auth-service/internal/repository/redis"
	"github.com/golang-jwt/jwt"
)

const (
	// salt набор случайных символов добавляемый к паролю
	salt            = "ljknsdfkgiovmsdlk&984kjsdlfj"
	accessTokenTTL  = 2 * time.Hour
	refreshTokenTTL = 7 * 24 * time.Hour
)

type accessTokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type refreshTokenClaims struct {
	jwt.StandardClaims
	UserId  int    `json:"user_id"`
	TokenId string `json:"token_id"`
}

type AuthService struct {
	repo      postgres.Authorization
	cacheRepo redis.Authorization
}

func NewAuthService(repo postgres.Authorization, cacheRepo redis.Authorization) *AuthService {
	return &AuthService{
		repo:      repo,
		cacheRepo: cacheRepo,
	}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	return 0, nil
}

func (s *AuthService) GenerateAccessToken(email, username, password_hash string) (string, error) {

	return "", nil
}

func (s *AuthService) UpdateAccessToken(userId int) (string, error) {
	return "", nil
}

func (s *AuthService) RevokingAllAccessToken(userId int) error {
	return nil
}

func (s *AuthService) GenerateRefreshToken(email, username, password_hash string) (string, error) {
	return "", nil
}

func (s *AuthService) ParseAccessToken(accessToken string) (int, error) {

	return 0, nil
}

func (s *AuthService) ParseRefreshToken(refreshToken string) (string, int, error) {

	return "", 0, nil
}

func (s *AuthService) IsRefreshTokenActive(tokenId string, userId int) (bool, error) {

	return true, nil
}

func (s *AuthService) DeleteRefreshToken(tokenId string, userId int) error {
	return nil
}

func (s *AuthService) DeleteAllRefreshToken(userId int) error {

	return nil
}
