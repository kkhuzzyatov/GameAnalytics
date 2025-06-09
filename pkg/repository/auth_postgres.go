package repository

import (
	"fmt"

	gameAnalytics "github.com/kkhuzzyatov/GameAnalytics"


	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user gameAnalytics.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUserByUsername(username string) (gameAnalytics.User, error) {
	var user gameAnalytics.User
	query := fmt.Sprintf("SELECT id, name, username, password_hash FROM %s WHERE username=$1", usersTable)
	err := r.db.Get(&user, query, username)

	return user, err
}
