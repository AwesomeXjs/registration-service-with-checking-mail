package model

import (
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserInfo struct {
	Email    string
	Password string
	Name     string
	Surname  string
	Role     string
}

type LoginInfo struct {
	Email    string
	Password string
}

type InfoToDb struct {
	ID           string       `db:"id" json:"id"`
	Email        string       `db:"email" json:"email"`
	HashPassword string       `db:"hash_password" json:"hash_password"`
	Role         string       `db:"role" json:"role"`
	CreatedAt    time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at" json:"updated_at"`
}

type AccessTokenInfo struct {
	ID   string
	Role string
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	UserId       string `json:"userId"`
}

type LoginResponse struct {
	UserID       string `db:"id"`
	HashPassword string `db:"hash_password"`
	Role         string `db:"role"`
}

type InfoToUserService struct {
	ID      string
	Name    string
	Surname string
}

type UserClaims struct {
	jwt.StandardClaims
	Role string
	ID   string
}
