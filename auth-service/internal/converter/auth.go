package converter

import "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"

// ToModelAuthResponse converts access and refresh tokens along with user ID
// into an AuthResponse model.
func ToModelAuthResponse(accessToken, refreshToken string, UserID int) *model.AuthResponse {
	return &model.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       int64(UserID),
	}
}

// ToModelAccessTokenInfo converts user ID and user information into an
// AccessTokenInfo model.
func ToModelAccessTokenInfo(userID int, userInfo *model.UserInfo) *model.AccessTokenInfo {
	return &model.AccessTokenInfo{
		ID:   userID,
		Role: userInfo.Role,
	}
}
