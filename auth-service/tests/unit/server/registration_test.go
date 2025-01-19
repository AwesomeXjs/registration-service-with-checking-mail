package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/app"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/grpc_server"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service/auth"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/validator"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/tests/unit/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestRegistration(t *testing.T) {
	t.Parallel()
	type IServiceMockFunc func(mc *minimock.Controller) auth.IServiceAuth
	logger.Init(logger.GetCore(logger.GetAtomicLevel(app.LogLevel)))

	type args struct {
		ctx context.Context
		req *authService.RegistrationRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, false, 8)
		name     = gofakeit.Name()
		surname  = gofakeit.LastName()
		role     = "user"
		userID   = 1

		accessToken  = gofakeit.UUID()
		refreshToken = gofakeit.UUID()

		req = &authService.RegistrationRequest{
			Email:    email,
			Password: password,
			Name:     name,
			Surname:  surname,
			Role:     role,
		}
		toService = converter.ToInfoFromProto(req)

		fromService = &model.AuthResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			UserID:       int64(userID),
		}

		res = converter.ToProtoFromRegResponse(fromService)

		serviceError = fmt.Errorf("service error")
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name         string
		args         args
		want         *authService.RegistrationResponse
		err          error
		IServiceMock IServiceMockFunc
	}{
		{name: "success",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			IServiceMock: func(mc *minimock.Controller) auth.IServiceAuth {
				service := mocks.NewIServiceAuthMock(mc)
				service.RegistrationMock.Expect(ctx, toService).Return(fromService, nil)
				return service
			},
		},

		{name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  fmt.Errorf("failed to registration: %s", serviceError.Error()),
			IServiceMock: func(mc *minimock.Controller) auth.IServiceAuth {
				service := mocks.NewIServiceAuthMock(mc)
				service.RegistrationMock.Expect(ctx, toService).Return(nil, serviceError)
				return service
			},
		},
		{name: "validate error case",
			args: args{
				ctx: ctx,
				req: &authService.RegistrationRequest{
					Email:    "123",
					Password: password,
					Name:     name,
					Surname:  surname,
					Role:     role,
				},
			},
			want: nil,
			err:  validator.NewValidationErrors("[\"invalid email\"]"),
			IServiceMock: func(mc *minimock.Controller) auth.IServiceAuth {
				service := mocks.NewIServiceAuthMock(mc)
				return service
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			IServiceMock := &service.Service{Auth: tt.IServiceMock(mc)}
			myController := grpc_server.New(IServiceMock)

			result, err := myController.Registration(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
