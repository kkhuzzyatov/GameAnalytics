package repository

import (
	gameAnalytics "github.com/kkhuzzyatov/GameAnalytics"


	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
)

type Authorization interface {
	CreateUser(user gameAnalytics.User) (int, error)
	GetUserByUsername(username string) (gameAnalytics.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}