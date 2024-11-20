package repository

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/db"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/redis"
)

// Repository provides database access for data operations.
type Repository struct {
	db          db.Client
	redisClient redis.IRedis
}

// New creates a new Repository instance with the given database client.
func New(db db.Client, redisClient redis.IRedis) IRepository {
	return &Repository{
		db:          db,
		redisClient: redisClient,
	}
}
