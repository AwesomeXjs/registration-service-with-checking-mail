package auth

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/jwt_manager"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/repository"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/db"
)

// IServiceAuth defines the interface for authentication-related operations.
type IServiceAuth interface {
	// Registration creates a new user account with the provided registration details.
	// It returns an AuthResponse containing tokens if successful, or an error if the registration fails.
	Registration(ctx context.Context, registrationRequest *model.UserInfo) (*model.AuthResponse, error)

	// Login authenticates a user using their login credentials.
	// It returns an AuthResponse containing tokens if successful, or an error if login fails.
	Login(ctx context.Context, loginRequest *model.LoginInfo) (*model.AuthResponse, error)

	// GetAccessToken generates a new access token using the provided refresh token.
	// It returns a NewPairTokens containing the new access and refresh tokens, or an error if the operation fails.
	GetAccessToken(ctx context.Context, refreshToken string) (*model.NewPairTokens, error)

	// ValidateToken checks the validity of the provided access token.
	// It returns true if the token is valid, along with any error encountered during validation.
	ValidateToken(ctx context.Context, accessToken string) (bool, error)

	// UpdatePassword changes the password for the user associated with the provided information.
	// It returns an error if the password update operation fails.
	UpdatePassword(ctx context.Context, updatePassInfo *model.UpdatePassInfo) error

	// ConfirmEmail sends a confirmation email to the user with the provided email address.
	// It returns an error if the email sending operation fails.
	ConfirmEmail(ctx context.Context, email string) error
}

// ServiceAuth provides authentication-related services, including repository interactions, JWT management, and transaction management.
type ServiceAuth struct {
	Repo       *repository.Repository
	AuthHelper jwt_manager.AuthHelper
	Tx         db.TxManager
}

// New creates a new instance of AuthService with the provided repository, JWT helper, and transaction manager.
func New(repo *repository.Repository, authHelper jwt_manager.AuthHelper, tx db.TxManager) IServiceAuth {
	return &ServiceAuth{
		Repo:       repo,
		AuthHelper: authHelper,
		Tx:         tx,
	}
}
