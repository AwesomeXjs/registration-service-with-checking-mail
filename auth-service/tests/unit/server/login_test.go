package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service/auth"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/validator"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/app"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/grpc_server"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/tests/unit/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	t.Parallel()
	type IServiceMockFunc func(mc *minimock.Controller) auth.IServiceAuth
	logger.Init(logger.GetCore(logger.GetAtomicLevel(app.LogLevel)))

	type args struct {
		ctx context.Context
		req *authService.LoginRequest
	}

	var (
		ctx      = context.Background()
		mc       = minimock.NewController(t)
		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, false, 8)

		accessToken  = gofakeit.UUID()
		refreshToken = gofakeit.UUID()
		userID       = 1

		req = &authService.LoginRequest{
			Email:    email,
			Password: password,
		}

		res = &authService.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			UserId:       int64(userID),
		}

		loginRequest = &model.LoginInfo{
			Email:    email,
			Password: password,
		}

		loginResponse = &model.AuthResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			UserID:       int64(userID),
		}

		serviceError = fmt.Errorf("service error")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name         string
		args         args
		want         *authService.LoginResponse
		err          error
		IServiceMock IServiceMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			IServiceMock: func(mc *minimock.Controller) auth.IServiceAuth {
				mock := mocks.NewIServiceAuthMock(mc)
				mock.LoginMock.Expect(ctx, loginRequest).Return(loginResponse, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceError,
			IServiceMock: func(mc *minimock.Controller) auth.IServiceAuth {
				mock := mocks.NewIServiceAuthMock(mc)
				// задаем поведение мока (все методы сервиса которые вызываются внутри ручки контроллера)
				mock.LoginMock.Expect(ctx, loginRequest).Return(nil, serviceError)
				return mock
			},
		},
		{
			name: "invalid email",
			args: args{
				ctx: ctx,
				req: &authService.LoginRequest{
					Email:    "123",
					Password: password,
				},
			},
			want: nil,
			err:  validator.NewValidationErrors("[\"invalid email\"]"),
			IServiceMock: func(mc *minimock.Controller) auth.IServiceAuth {
				mock := mocks.NewIServiceAuthMock(mc)
				return mock
			},
		},
		{
			name: "invalid password",
			args: args{
				ctx: ctx,
				req: &authService.LoginRequest{
					Email:    email,
					Password: "123",
				},
			},
			want: nil,
			err:  validator.NewValidationErrors("[\"password length must be between 5 and 20 characters\"]"),
			IServiceMock: func(mc *minimock.Controller) auth.IServiceAuth {
				mock := mocks.NewIServiceAuthMock(mc)
				return mock
			},
		},
		{
			name: "invalid password and email",
			args: args{
				ctx: ctx,
				req: &authService.LoginRequest{
					Email:    "123",
					Password: "123",
				},
			},
			want: nil,
			err:  validator.NewValidationErrors("[\"invalid email\"]", "[\"password length must be between 5 and 20 characters\"]"),
			IServiceMock: func(mc *minimock.Controller) auth.IServiceAuth {
				mock := mocks.NewIServiceAuthMock(mc)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			IServiceMock := &service.Service{Auth: tt.IServiceMock(mc)}
			myController := grpc_server.New(IServiceMock)

			result, err := myController.Login(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
