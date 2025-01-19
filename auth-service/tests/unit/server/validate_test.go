package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/app"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/grpc_server"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service/auth"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/tests/unit/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	type IServiceMockFunc func(mc *minimock.Controller) auth.IServiceAuth
	logger.Init(logger.GetCore(logger.GetAtomicLevel(app.LogLevel)))

	type args struct {
		ctx context.Context
		req *authService.ValidateTokenRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		accessToken = gofakeit.UUID()

		req = &authService.ValidateTokenRequest{
			AccessToken: accessToken,
		}

		serviceError = fmt.Errorf("some error")
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
				mock := mocks.NewIServiceAuthMock(mc)
				mock.ValidateTokenMock.Expect(ctx, accessToken).Return(true, nil)
				return mock
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
				mock := mocks.NewIServiceAuthMock(mc)
				mock.ValidateTokenMock.Expect(ctx, accessToken).Return(false, serviceError)
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

			result, err := myController.ValidateToken(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
