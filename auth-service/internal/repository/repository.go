package repository

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/model"
)

// IRepository defines the methods for user-related database operations.
// It abstracts the interactions with the data layer for user authentication and management.
type IRepository interface {
	// Registration saves a new user's registration information to the database.
	// It returns the user's ID upon successful registration or an error if the operation fails.
	Registration(ctx context.Context, registrationRequest *model.InfoToDb) (string, error)

	// Login retrieves a user's login information based on their email.
	// It returns a LoginResponse containing the user's credentials or an error if the operation fails.
	Login(ctx context.Context, email string) (*model.LoginResponse, error)

	// GetAccessToken fetches access token information for a user identified by their user ID.
	// It returns the AccessTokenInfo struct containing relevant token details or an error if the operation fails.
	GetAccessToken(ctx context.Context, userID string) (*model.AccessTokenInfo, error)

	// UpdatePassword updates the user's password in the database.
	// It takes a UpdatePassDb struct with the user's email and new hashed password,
	// and returns an error if the operation fails.
	UpdatePassword(ctx context.Context, updatePassDb *model.UpdatePassDb) error
}
