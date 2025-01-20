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

func TestGetAccessToken(t *testing.T) {
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

		accessToken  = gofakeit.UUID()
		refreshToken = gofakeit.UUID()
		userID       = 1
		role         = "user"

		res = &model.NewPairTokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}
		tokenInfo = &model.AccessTokenInfo{
			ID:   userID,
			Role: role,
		}

		userClaims = &model.UserClaims{
			ID:   userID,
			Role: role,
		}

		someError = fmt.Errorf("some error")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name     string
		args     args
		want     *model.NewPairTokens
		err      error
		repoMock IRepositoryMockFunc
		jwtMock  JWTHelperMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx:          ctx,
				refreshToken: refreshToken,
			},
			want: res,
			err:  nil,
			repoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				repo := mocks.NewIRepositoryAuthMock(mc)
				repo.GetAccessTokenMock.Expect(ctx, userID).Return(tokenInfo, nil)
				return repo
			},
			jwtMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				jwt := mocks.NewAuthHelperMock(mc)
				jwt.VerifyTokenMock.Expect(refreshToken).Return(userClaims, nil)
				jwt.GenerateAccessTokenMock.Expect(tokenInfo).Return(accessToken, nil)
				jwt.GenerateRefreshTokenMock.Expect(userID).Return(refreshToken, nil)
				return jwt
			},
		},
		{
			name: "repository error",
			args: args{
				ctx:          ctx,
				refreshToken: refreshToken,
			},
			want: nil,
			err:  fmt.Errorf("failed to get access token: %s", someError.Error()),
			repoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				repo := mocks.NewIRepositoryAuthMock(mc)
				repo.GetAccessTokenMock.Expect(ctx, userID).Return(tokenInfo, someError)
				return repo
			},
			jwtMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				jwt := mocks.NewAuthHelperMock(mc)
				jwt.VerifyTokenMock.Expect(refreshToken).Return(userClaims, nil)
				return jwt
			},
		},
		{
			name: "failed to verify token case",
			args: args{
				ctx:          ctx,
				refreshToken: refreshToken,
			},
			want: nil,
			err:  fmt.Errorf("failed to verify token: %s", someError.Error()),
			repoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				repo := mocks.NewIRepositoryAuthMock(mc)
				return repo
			},
			jwtMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				jwt := mocks.NewAuthHelperMock(mc)
				jwt.VerifyTokenMock.Expect(refreshToken).Return(userClaims, someError)
				return jwt
			},
		},
		{
			name: "failed to generate access token case",
			args: args{
				ctx:          ctx,
				refreshToken: refreshToken,
			},
			want: nil,
			err:  fmt.Errorf("failed to generate access token: %s", someError.Error()),
			repoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				repo := mocks.NewIRepositoryAuthMock(mc)
				repo.GetAccessTokenMock.Expect(ctx, userID).Return(tokenInfo, nil)
				return repo
			},
			jwtMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				jwt := mocks.NewAuthHelperMock(mc)
				jwt.VerifyTokenMock.Expect(refreshToken).Return(userClaims, nil)
				jwt.GenerateAccessTokenMock.Expect(tokenInfo).Return(accessToken, someError)
				return jwt
			},
		},
		{
			name: "failed to generate refresh token case",
			args: args{
				ctx:          ctx,
				refreshToken: refreshToken,
			},
			want: nil,
			err:  fmt.Errorf("failed to generate refresh token: %s", someError.Error()),
			repoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				repo := mocks.NewIRepositoryAuthMock(mc)
				repo.GetAccessTokenMock.Expect(ctx, userID).Return(tokenInfo, nil)
				return repo
			},
			jwtMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				jwt := mocks.NewAuthHelperMock(mc)
				jwt.VerifyTokenMock.Expect(refreshToken).Return(userClaims, nil)
				jwt.GenerateAccessTokenMock.Expect(tokenInfo).Return(accessToken, nil)
				jwt.GenerateRefreshTokenMock.Expect(userID).Return(refreshToken, someError)
				return jwt
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			IRepositoryMock := &repository.Repository{Auth: tt.repoMock(mc)}
			service := serviceAuth.ServiceAuth{Repo: IRepositoryMock, AuthHelper: tt.jwtMock(mc)}

			result, err := service.GetAccessToken(tt.args.ctx, tt.args.refreshToken)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
