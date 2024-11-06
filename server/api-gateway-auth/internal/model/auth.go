package model

type LoginRequest struct {
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required,stringlength(5|20)"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	UserID       string `json:"userId"`
}

type GetAccessTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type RegistrationRequest struct {
	Email    string `json:"email" valid:"required,email,type(string)"`
	Password string `json:"password" valid:"required,type(string)"`
	Name     string `json:"name" valid:"required,type(string),stringlength(3|64)"`
	Role     string `json:"role" valid:"required,type(string),role_enum"`
	Surname  string `json:"surname" valid:"required,type(string)"`
}

type RegistrationResponse struct {
	AccessToken string `json:"accessToken"`
	UserID      string `json:"userId"`
}

type UpdatePasswordRequest struct {
	Email       string `json:"email" valid:"required,email,type(string)"`
	NewPassword string `json:"newPassword" valid:"required,type(string),stringlength(5|20)"`
}
