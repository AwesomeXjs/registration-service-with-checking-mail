package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/app"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/grpc_server"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service/auth"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/tests/unit/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestGetAccessToken(t *testing.T) {
	t.Parallel()

	type IServiceMockFunc func(mc *minimock.Controller) auth.IServiceAuth
	logger.Init(logger.GetCore(logger.GetAtomicLevel(app.LogLevel)))

	type args struct {
		ctx context.Context
		req *authService.GetAccessTokenRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		refreshToken = gofakeit.UUID()
		accessToken  = gofakeit.UUID()

		req = &authService.GetAccessTokenRequest{
			RefreshToken: refreshToken,
		}

		res = &authService.GetAccessTokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		serviceError = fmt.Errorf("some error")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name         string
		args         args
		want         *authService.GetAccessTokenResponse
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
				mock.GetAccessTokenMock.Expect(ctx, req.RefreshToken).Return(&model.NewPairTokens{
					AccessToken:  accessToken,
					RefreshToken: refreshToken,
				}, nil)
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
				mock.GetAccessTokenMock.Expect(ctx, req.RefreshToken).Return(nil, serviceError)
				return mock
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			IServiceMock := &service.Service{Auth: tt.IServiceMock(mc)}
			controller := grpc_server.New(IServiceMock)

			result, err := controller.GetAccessToken(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
