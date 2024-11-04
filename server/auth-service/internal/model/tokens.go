package model

import "github.com/golang-jwt/jwt"

// AccessTokenInfo holds information extracted from the access token,
// including the user's ID and role. This is used for authorization checks.
type AccessTokenInfo struct {
	ID   string `db:"id"`
	Role string `db:"role"`
}

// AuthResponse represents the structure of an authentication response.
// It contains the access token, refresh token, and the user's ID.
// This is typically returned after successful login or registration.
type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	UserID       string `json:"userId"`
}

// UserClaims extends the standard JWT claims to include custom fields
// such as the user's role and ID. These claims are used to manage
// user identity and permissions.
type UserClaims struct {
	jwt.StandardClaims
	Role string
	ID   string
}

// NewPairTokens holds a new set of access and refresh tokens.
// This is used when refreshing tokens to maintain user session validity.
type NewPairTokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
