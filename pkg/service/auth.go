package service

import (
	gameAnalytics "github.com/kkhuzzyatov/GameAnalytics"

	"errors"

	"github.com/kkhuzzyatov/GameAnalytics/pkg/repository"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	signingKey = "F;f4902k5fkK)(#NF&5g58hfa;f]f"
	// tokenTTL   = 24 * time.Hour
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user gameAnalytics.User) (int, error) {
	hashedPassword, err := s.generatePasswordHash(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hashedPassword
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			// ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid siging method")
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

func (s *AuthService) generatePasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (s *AuthService) ValidatePassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}
