package repository

import "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/clients/db"

type Repository struct {
	db db.Client
}

func New(db db.Client) *Repository {
	return &Repository{
		db: db,
	}
}
