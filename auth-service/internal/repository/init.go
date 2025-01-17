package repository

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/redis"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/repository/auth"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/repository/events"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/db"
)

// Repository provides database access for data operations.
type Repository struct {
	Auth  auth.IRepositoryAuth
	Event events.IEventRepository
}

// New creates a new Repository instance with the given database client.
func New(db db.Client, redisClient redis.IRedis) *Repository {
	return &Repository{
		Auth:  auth.New(db, redisClient),
		Event: events.New(db),
	}
}
