package model

import "github.com/golang-jwt/jwt"

type AccessTokenInfo struct {
	ID   string `db:"id"`
	Role string `db:"role"`
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	UserId       string `json:"userId"`
}

type UserClaims struct {
	jwt.StandardClaims
	Role string
	ID   string
}

type NewPairTokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
