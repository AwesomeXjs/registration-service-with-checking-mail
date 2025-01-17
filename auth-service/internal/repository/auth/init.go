package auth

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/redis"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/db"
)

// IRepositoryAuth defines the interface for authentication-related database and Redis operations.
type IRepositoryAuth interface {
	// Registration saves a new user's registration information to the database.
	// It returns the user's ID upon successful registration or an error if the operation fails.
	Registration(ctx context.Context, registrationRequest *model.InfoToDb) (int, error)

	// Login retrieves a user's login information based on their email.
	// It returns a LoginResponse containing the user's credentials or an error if the operation fails.
	Login(ctx context.Context, email string) (*model.LoginResponse, error)

	// GetAccessToken fetches access token information for a user identified by their user ID.
	// It returns the AccessTokenInfo struct containing relevant token details or an error if the operation fails.
	GetAccessToken(ctx context.Context, userID int) (*model.AccessTokenInfo, error)

	// UpdatePassword updates the user's password in the database.
	// It takes a UpdatePassDb struct with the user's email and new hashed password,
	// and returns an error if the operation fails.
	UpdatePassword(ctx context.Context, updatePassDb *model.UpdatePassDb) error

	// ConfirmEmail sends a confirmation email to the user with the provided email address.
	// It returns an error if the email sending operation fails.
	ConfirmEmail(ctx context.Context, email string) error
}

// RepositoryAuth provides methods for authentication-related database and Redis operations.
type RepositoryAuth struct {
	db          db.Client    // Database client for executing queries.
	redisClient redis.IRedis // Redis client for cache operations.
}

// New creates a new instance of AuthRepository with provided database and Redis clients.
func New(db db.Client, redisClient redis.IRedis) IRepositoryAuth {
	return &RepositoryAuth{
		db:          db,
		redisClient: redisClient,
	}
}
