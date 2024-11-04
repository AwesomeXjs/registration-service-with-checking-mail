package authHelper

import "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/model"

type AuthHelper interface {
	GenerateAccessToken(info *model.AccessTokenInfo) (string, error)
	GenerateRefreshToken(userID string) (string, error)
	VerifyToken(token string) (*model.UserClaims, error)
	HashPassword(password string) (string, error)
	ValidatePassword(hashedPassword, candidatePassword string) bool
}
