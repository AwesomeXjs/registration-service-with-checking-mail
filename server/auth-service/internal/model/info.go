package model

import (
	"database/sql"
	"time"
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

type UpdatePassInfo struct {
	Email       string
	NewPassword string
}

type UpdatePassDb struct {
	Email        string `db:"email"`
	HashPassword string `db:"hash_password"`
}
