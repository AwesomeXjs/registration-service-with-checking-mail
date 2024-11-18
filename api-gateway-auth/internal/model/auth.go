package model

// LoginRequest represents the structure for a login request, including email and password fields.
// Email must be a valid email address and is required.
// Password is required and must be a string between 5 and 20 characters in length.
type LoginRequest struct {
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required,stringlength(5|20)"`
}

// LoginResponse represents the structure for a successful login response,
// containing the generated access token and the user ID.
type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	UserID      string `json:"userId"`
}

// RegistrationRequest represents the structure for a registration request,
// including the user's email, password, name, role, and surname.
// Email must be a valid email address and is required.
// Password must be a string and is required.
// Name and surname must be strings, with name having a length between 3 and 64 characters.
// Role must be a string and match a predefined enum of roles.
type RegistrationRequest struct {
	Email    string `json:"email" valid:"required,email,type(string)"`
	Password string `json:"password" valid:"required,type(string)"`
	Name     string `json:"name" valid:"required,type(string),stringlength(3|64)"`
	Role     string `json:"role" valid:"required,type(string),role_enum"`
	Surname  string `json:"surname" valid:"required,type(string)"`
}

// RegistrationResponse represents the structure for a successful registration response,
// containing the generated access token and the user ID.
type RegistrationResponse struct {
	AccessToken string `json:"accessToken"`
	UserID      string `json:"userId"`
}

// UpdatePasswordRequest represents the structure for a password update request,
// including the user's email and new password.
// Email must be a valid email address and is required.
// NewPassword must be a string between 5 and 20 characters and is required.
type UpdatePasswordRequest struct {
	Email       string `json:"email" valid:"required,email,type(string)"`
	NewPassword string `json:"newPassword" valid:"required,type(string),stringlength(5|20)"`
}
