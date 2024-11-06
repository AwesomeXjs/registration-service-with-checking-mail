package model

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	UserID       string `json:"userId"`
}

type GetAccessTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type GetAccessTokenResponse struct {
	AccessToken string `json:"accessToken"`
}
