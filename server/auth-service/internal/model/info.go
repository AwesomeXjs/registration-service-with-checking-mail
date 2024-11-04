package model

import (
	"database/sql"
	"time"
)

// UserInfo holds the basic user information needed for registration and authentication.
// This includes email, password, name, surname, and the user's role.
type UserInfo struct {
	Email    string
	Password string
	Name     string
	Surname  string
	Role     string
}

// LoginInfo is a simplified structure containing only the user's email and password.
// It is used specifically during the login process.
type LoginInfo struct {
	Email    string
	Password string
}

// InfoToDb represents the structure used for interacting with the database for user records.
// It includes fields for the user's ID, Email, HashPassword, Role, and timestamps for record creation and updates.
type InfoToDb struct {
	ID           string       `db:"id" json:"id"`
	Email        string       `db:"email" json:"email"`
	HashPassword string       `db:"hash_password" json:"hash_password"`
	Role         string       `db:"role" json:"role"`
	CreatedAt    time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at" json:"updated_at"`
}

// LoginResponse is used to retrieve essential authentication details from the database.
// It contains the user's ID, hashed password, and role information.
type LoginResponse struct {
	UserID       string `db:"id"`
	HashPassword string `db:"hash_password"`
	Role         string `db:"role"`
}

// InfoToUserService is a simplified structure passed to the User Service,
// containing the user's ID, name, and surname.
type InfoToUserService struct {
	ID      string
	Name    string
	Surname string
}

// UpdatePassInfo holds the email and new password for a user who wants to update their password.
// This structure is used during password change requests.
type UpdatePassInfo struct {
	Email       string
	NewPassword string
}

// UpdatePassDb defines the structure for updating a user's password in the database.
// It includes the user's email, new hashed password, and the updated timestamp.
type UpdatePassDb struct {
	Email        string    `db:"email"`
	HashPassword string    `db:"hash_password"`
	UpdatedAt    time.Time `db:"updated_at"`
}
