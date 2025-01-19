package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/app"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/grpc_server"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service/auth"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/validator"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/tests/unit/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestUpdatePassword(t *testing.T) {
	t.Parallel()

	type IServiceMockFunc func(mc *minimock.Controller) auth.IServiceAuth
	logger.Init(logger.GetCore(logger.GetAtomicLevel(app.LogLevel)))

	type args struct {
		ctx context.Context
		req *authService.UpdatePasswordRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, false, 8)

		req = &authService.UpdatePasswordRequest{
			Email:       email,
			NewPassword: password,
		}

		validationError = validator.NewValidationErrors("[\"invalid email\"]")
		serviceError    = fmt.Errorf("some error")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name         string
		args         args
		want         *emptypb.Empty
		err          error
		IServiceMock IServiceMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  nil,
			IServiceMock: func(mc *minimock.Controller) auth.IServiceAuth {
				service := mocks.NewIServiceAuthMock(mc)
				service.UpdatePasswordMock.Expect(ctx, converter.ToUpdatePassFromProto(req)).Return(nil)
				return service
			},
		},
		{
			name: "validate error",
			args: args{
				ctx: ctx,
				req: &authService.UpdatePasswordRequest{
					Email:       "123",
					NewPassword: password,
				},
			},
			want: nil,
			err:  validationError,
			IServiceMock: func(mc *minimock.Controller) auth.IServiceAuth {
				service := mocks.NewIServiceAuthMock(mc)
				return service
			},
		},
		{
			name: "service error",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceError,
			IServiceMock: func(mc *minimock.Controller) auth.IServiceAuth {
				service := mocks.NewIServiceAuthMock(mc)
				service.UpdatePasswordMock.Expect(ctx, converter.ToUpdatePassFromProto(req)).Return(serviceError)
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

			result, err := myController.UpdatePassword(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}

}
