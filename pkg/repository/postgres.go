package repository

import (

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(DSN string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", DSN)
	if err != nil {
		logrus.Fatalf("failed to connect to database: %s", err.Error())
		return nil, err
	}

	if err := db.Ping(); err != nil {
		logrus.Fatalf("failed to connect to database: %s", err.Error())
		return nil, err
	}

	return db, nil
}
