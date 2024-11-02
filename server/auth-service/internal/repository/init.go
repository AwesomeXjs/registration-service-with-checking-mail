package repository

import "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/clients/db"

// Repository provides database access for data operations.
type Repository struct {
	db db.Client
}

// New creates a new Repository instance with the given database client.
func New(db db.Client) *Repository {
	return &Repository{
		db: db,
	}
}
