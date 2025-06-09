package service

import (
	gameAnalytics "github.com/kkhuzzyatov/GameAnalytics"

	"github.com/kkhuzzyatov/GameAnalytics/pkg/repository"
)

type Authorization interface {
	CreateUser(user gameAnalytics.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
