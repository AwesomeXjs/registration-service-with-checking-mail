package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/app"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/jwt_manager"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/repository"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/repository/auth"
	serviceAuth "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service/auth"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/tests/unit/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	t.Parallel()
	logger.Init(logger.GetCore(logger.GetAtomicLevel(app.LogLevel)))

	type IRepositoryMockFunc func(mc *minimock.Controller) auth.IRepositoryAuth
	type JWTHelperMockFunc func(mc *minimock.Controller) jwt_manager.AuthHelper

	type args struct {
		ctx          context.Context
		refreshToken string
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, false, 8)

		accessToken  = gofakeit.UUID()
		refreshToken = gofakeit.UUID()
		userID       = 1
		hashPassword = gofakeit.UUID()

		req = &model.LoginInfo{
			Email:    email,
			Password: password,
		}
		loginResponse = &model.LoginResponse{
			UserID:       userID,
			HashPassword: hashPassword,
			Role:         "user",
		}
		accessTokenInfo = &model.AccessTokenInfo{
			ID:   userID,
			Role: "user",
		}

		res = &model.AuthResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			UserID:       int64(userID),
		}

		someError = fmt.Errorf("some error")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name     string
		args     args
		want     *model.AuthResponse
		err      error
		repoMock IRepositoryMockFunc
		jwtMock  JWTHelperMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx:          ctx,
				refreshToken: refreshToken,
			},
			want: res,
			err:  nil,
			repoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				repo := mocks.NewIRepositoryAuthMock(mc)
				repo.LoginMock.Expect(ctx, email).Return(loginResponse, nil)
				return repo
			},
			jwtMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				jwt := mocks.NewAuthHelperMock(mc)
				jwt.ValidatePasswordMock.Expect(hashPassword, password).Return(true)
				jwt.GenerateAccessTokenMock.Expect(accessTokenInfo).Return(accessToken, nil)
				jwt.GenerateRefreshTokenMock.Expect(userID).Return(refreshToken, nil)
				return jwt
			},
		},
		{
			name: "repository error case",
			args: args{
				ctx:          ctx,
				refreshToken: refreshToken,
			},
			want: nil,
			err:  someError,
			repoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				repo := mocks.NewIRepositoryAuthMock(mc)
				repo.LoginMock.Expect(ctx, email).Return(loginResponse, someError)
				return repo
			},
			jwtMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				jwt := mocks.NewAuthHelperMock(mc)
				return jwt
			},
		},
		{
			name: "validate password error case",
			args: args{
				ctx:          ctx,
				refreshToken: refreshToken,
			},
			want: nil,
			err:  fmt.Errorf("invalid password"),
			repoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				repo := mocks.NewIRepositoryAuthMock(mc)
				repo.LoginMock.Expect(ctx, email).Return(loginResponse, nil)
				return repo
			},
			jwtMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				jwt := mocks.NewAuthHelperMock(mc)
				jwt.ValidatePasswordMock.Expect(hashPassword, password).Return(false)
				return jwt
			},
		},
		{
			name: "generate access token error case",
			args: args{
				ctx:          ctx,
				refreshToken: refreshToken,
			},
			want: nil,
			err:  fmt.Errorf("failed to generate access token: %s", someError.Error()),
			repoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				repo := mocks.NewIRepositoryAuthMock(mc)
				repo.LoginMock.Expect(ctx, email).Return(loginResponse, nil)
				return repo
			},
			jwtMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				jwt := mocks.NewAuthHelperMock(mc)
				jwt.ValidatePasswordMock.Expect(hashPassword, password).Return(true)
				jwt.GenerateAccessTokenMock.Expect(accessTokenInfo).Return(accessToken, someError)
				return jwt
			},
		},
		{
			name: "generate refresh token error case",
			args: args{
				ctx:          ctx,
				refreshToken: refreshToken,
			},
			want: nil,
			err:  fmt.Errorf("failed to generate refresh token: %s", someError.Error()),
			repoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				repo := mocks.NewIRepositoryAuthMock(mc)
				repo.LoginMock.Expect(ctx, email).Return(loginResponse, nil)
				return repo
			},
			jwtMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				jwt := mocks.NewAuthHelperMock(mc)
				jwt.ValidatePasswordMock.Expect(hashPassword, password).Return(true)
				jwt.GenerateAccessTokenMock.Expect(accessTokenInfo).Return(accessToken, nil)
				jwt.GenerateRefreshTokenMock.Expect(userID).Return(refreshToken, someError)
				return jwt
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			IRepositoryMock := &repository.Repository{Auth: tt.repoMock(mc)}
			service := serviceAuth.ServiceAuth{Repo: IRepositoryMock, AuthHelper: tt.jwtMock(mc)}

			result, err := service.Login(tt.args.ctx, req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}

}
