package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/repository"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/service"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/auth_helper"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/tests/unit/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	t.Parallel()
	level := "info"
	logger.Init(logger.GetCore(logger.GetAtomicLevel(&level)))

	type IRepositoryMockFunc func(mc *minimock.Controller) repository.IRepository
	type AuthHelperMockFunc func(mc *minimock.Controller) auth_helper.AuthHelper

	type args struct {
		ctx context.Context
		req *model.LoginInfo
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		email        = gofakeit.Email()
		password     = gofakeit.Password(true, true, true, true, false, 8)
		hashPassword = gofakeit.UUID()

		accessToken  = gofakeit.UUID()
		refreshToken = gofakeit.UUID()
		userID       = gofakeit.UUID()
		role         = "user"

		req = &model.LoginInfo{
			Email:    email,
			Password: password,
		}

		res = &model.AuthResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			UserID:       userID,
		}

		loginResponse = &model.LoginResponse{
			UserID:       userID,
			HashPassword: hashPassword,
			Role:         role,
		}

		repositoryError = fmt.Errorf("repository error")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string              // название теста
		args            args                // аргументы которые передаем в ручку Login
		want            *model.AuthResponse // что хотим получить из ручки Login
		err             error
		IRepositoryMock IRepositoryMockFunc // функция которая возвращает замоканый сервис с нужным поведением
		AuthHelperMock  AuthHelperMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			IRepositoryMock: func(mc *minimock.Controller) repository.IRepository {
				mock := mocks.NewIRepositoryMock(mc)
				// задаем поведение мока (все методы сервиса которые вызываются внутри ручки контроллера)
				mock.LoginMock.Expect(ctx, email).Return(loginResponse, nil)
				return mock
			},
			AuthHelperMock: func(mc *minimock.Controller) auth_helper.AuthHelper {
				mock := mocks.NewAuthHelperMock(mc)
				mock.ValidatePasswordMock.Expect(hashPassword, password).Return(true)
				mock.GenerateAccessTokenMock.Expect(&model.AccessTokenInfo{
					ID:   userID,
					Role: role,
				}).Return(accessToken, nil)
				mock.GenerateRefreshTokenMock.Expect(userID).Return(refreshToken, nil)
				return mock
			},
		},
		{
			name: "repository error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  repositoryError,
			IRepositoryMock: func(mc *minimock.Controller) repository.IRepository {
				mock := mocks.NewIRepositoryMock(mc)
				// задаем поведение мока (все методы сервиса которые вызываются внутри ручки контроллера)
				mock.LoginMock.Expect(ctx, email).Return(nil, repositoryError)
				return mock
			},
			AuthHelperMock: func(mc *minimock.Controller) auth_helper.AuthHelper {
				mock := mocks.NewAuthHelperMock(mc)
				return mock
			},
		},
		{
			name: "validate password error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  fmt.Errorf("invalid password"),
			IRepositoryMock: func(mc *minimock.Controller) repository.IRepository {
				mock := mocks.NewIRepositoryMock(mc)
				// задаем поведение мока (все методы сервиса которые вызываются внутри ручки контроллера)
				mock.LoginMock.Expect(ctx, email).Return(loginResponse, nil)
				return mock
			},
			AuthHelperMock: func(mc *minimock.Controller) auth_helper.AuthHelper {
				mock := mocks.NewAuthHelperMock(mc)
				mock.ValidatePasswordMock.Expect(hashPassword, password).Return(false)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			IRepoMock := tt.IRepositoryMock(mc)
			AuthHelperMock := tt.AuthHelperMock(mc)

			myService := service.New(IRepoMock, AuthHelperMock)

			result, err := myService.Login(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}