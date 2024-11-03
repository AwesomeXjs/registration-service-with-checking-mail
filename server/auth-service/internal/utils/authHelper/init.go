package authHelper

import (
	"fmt"
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type AuthClient struct {
	secretKey            []byte
	refreshTokenDuration time.Duration
	accessTokenDuration  time.Duration
}

func New(secretKey []byte, refreshTokenDuration time.Duration, accessTokenDuration time.Duration) AuthHelper {
	return &AuthClient{
		secretKey:            secretKey,
		refreshTokenDuration: refreshTokenDuration,
		accessTokenDuration:  accessTokenDuration,
	}
}

func (a *AuthClient) GenerateAccessToken(info *model.InfoToDb) (string, error) {
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

func (a *AuthClient) GenerateRefreshToken(info *model.InfoToDb) (string, error) {
	claims := model.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(a.refreshTokenDuration).Unix(),
		},
		ID: info.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.secretKey)
}

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
	if !ok {
		return nil, err
	}
	return claims, nil
}

func (a *AuthClient) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("failed to hash password", zap.Error(err))
		return "", err
	}
	return string(hashedPassword), nil
}

func (a *AuthClient) ValidatePassword(hashedPassword, candidatePassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
	return err == nil
}
