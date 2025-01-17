package jwt_manager

import (
	"fmt"
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// AuthClient implements the AuthHelper interface and provides methods for
// managing authentication tasks, including token generation and password handling.
type AuthClient struct {
	secretKey            []byte        // Secret key used for signing tokens
	refreshTokenDuration time.Duration // Duration for which the refresh token is valid
	accessTokenDuration  time.Duration // Duration for which the access token is valid
}

// New creates a new AuthClient instance with the provided secret key and token durations.
// It returns an AuthHelper interface implementation for managing authentication tasks.
func New(secretKey []byte, refreshTokenDuration time.Duration, accessTokenDuration time.Duration) AuthHelper {
	return &AuthClient{
		secretKey:            secretKey,
		refreshTokenDuration: refreshTokenDuration,
		accessTokenDuration:  accessTokenDuration,
	}
}

// GenerateAccessToken creates a new access token for a user based on the provided access token information.
// It sets the expiration time for the token and returns the signed token as a string.
func (a *AuthClient) GenerateAccessToken(info *model.AccessTokenInfo) (string, error) {
	claims := model.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(a.accessTokenDuration).Unix(),
		},
		ID:   info.ID,
		Role: info.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.secretKey)
}

// GenerateRefreshToken creates a new refresh token for a user identified by their user ID.
// The token will have an expiration time set according to the specified duration.
func (a *AuthClient) GenerateRefreshToken(userID int) (string, error) {
	claims := model.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(a.refreshTokenDuration).Unix(),
		},
		ID: userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.secretKey)
}

// VerifyToken checks the validity of the provided JWT token. It parses the token and extracts the claims,
// returning them if the token is valid. An error is returned if the token is invalid or expired.
func (a *AuthClient) VerifyToken(token string) (*model.UserClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &model.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := tokenClaims.Claims.(*model.UserClaims)
	if !ok || !tokenClaims.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

// HashPassword generates a hashed version of the provided password using bcrypt.
// It returns the hashed password as a string or an error if the hashing fails.
func (a *AuthClient) HashPassword(password string) (string, error) {
	const mark = "JwtManager.Init.HashPassword"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("failed to hash password", mark, zap.Error(err))
		return "", err
	}
	return string(hashedPassword), nil
}

// ValidatePassword compares a hashed password with a candidate password to verify their equality.
// It returns true if they match, or false otherwise.
func (a *AuthClient) ValidatePassword(hashedPassword, candidatePassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
	return err == nil
}
