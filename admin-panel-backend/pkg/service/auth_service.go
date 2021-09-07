package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/Hudayberdyyev/admin-panel-backend/models"
	"github.com/Hudayberdyyev/admin-panel-backend/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	hashKey = "1jan@%!)(^"
	signingKey = "2JaaBaaZABA!@#$%^&*()_)"
	tokenTTL = 1 * time.Hour
)

type AuthService struct {
	repos repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{ repos: repos }
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repos.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	userId, err := s.repos.GetUserId(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) Parse(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, errors.New("invalid token method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(hashKey)))
}