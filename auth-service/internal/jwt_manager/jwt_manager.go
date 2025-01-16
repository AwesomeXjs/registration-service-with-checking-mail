package jwt_manager

import "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"

// AuthHelper defines the methods required for authentication tasks,
// including token generation, verification, and password management.
type AuthHelper interface {
	// GenerateAccessToken creates a new access token for the user,
	// using the provided access token information. Returns the token
	// as a string or an error if the operation fails.
	GenerateAccessToken(info *model.AccessTokenInfo) (string, error)

	// GenerateRefreshToken creates a new refresh token for a user,
	// identified by their user ID. Returns the refresh token as a
	// string or an error if the operation fails.
	GenerateRefreshToken(userID int) (string, error)

	// VerifyToken checks the validity of the provided token and
	// returns the associated user claims if the token is valid.
	// If the token is invalid or expired, an error is returned.
	VerifyToken(token string) (*model.UserClaims, error)

	// HashPassword takes a plain-text password and returns the
	// hashed version of it. This is used for securely storing
	// user passwords in the database. It may return an error
	// if the hashing process fails.
	HashPassword(password string) (string, error)

	// ValidatePassword compares a hashed password with a candidate
	// password to determine if they match. Returns true if they
	// are the same, or false otherwise.
	ValidatePassword(hashedPassword, candidatePassword string) bool
}
