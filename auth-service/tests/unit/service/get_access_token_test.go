package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/kafka"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/jwt_manager"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/repository"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/tests/unit/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestGetAccessToken(t *testing.T) {
	t.Parallel()
	level := info
	logger.Init(logger.GetCore(logger.GetAtomicLevel(&level)))

	type IRepositoryMockFunc func(mc *minimock.Controller) repository.IRepository
	type AuthHelperMockFunc func(mc *minimock.Controller) jwt_manager.AuthHelper
	type IProducerMockFunc func(mc *minimock.Controller) kafka.IProducer

	type args struct {
		ctx context.Context
		req string
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		refreshToken = gofakeit.UUID()
		accessToken  = gofakeit.UUID()
		userID       = 1

		res = &model.NewPairTokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		repositoryError  = fmt.Errorf("failed to login {\"error\": \"repository error\"}")
		verifyTokenError = fmt.Errorf("failed to verify token {\"error\": \"verify token error\"}")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name              string
		args              args
		want              *model.NewPairTokens
		err               error
		IRepositoryMock   IRepositoryMockFunc
		AuthHelperMock    AuthHelperMockFunc
		IProducerMockFunc IProducerMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: refreshToken,
			},
			want: res,
			err:  nil,
			IRepositoryMock: func(mc *minimock.Controller) repository.IRepository {
				mock := mocks.NewIRepositoryMock(mc)
				// задаем поведение мока (все методы сервиса которые вызываются внутри ручки контроллера)
				mock.GetAccessTokenMock.Expect(ctx, userID).Return(&model.AccessTokenInfo{
					ID:   userID,
					Role: "user",
				}, nil)
				return mock
			},
			AuthHelperMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				mock := mocks.NewAuthHelperMock(mc)
				mock.VerifyTokenMock.Expect(refreshToken).Return(&model.UserClaims{
					ID:   userID,
					Role: "user",
				}, nil)
				mock.GenerateAccessTokenMock.Expect(&model.AccessTokenInfo{
					ID:   userID,
					Role: "user",
				}).Return(accessToken, nil)
				mock.GenerateRefreshTokenMock.Expect(userID).Return(refreshToken, nil)
				return mock
			},
			IProducerMockFunc: func(mc *minimock.Controller) kafka.IProducer {
				mock := mocks.NewIProducerMock(mc)
				return mock
			},
		},
		{
			name: "repository error case",
			args: args{
				ctx: ctx,
				req: refreshToken,
			},
			want: nil,
			err:  fmt.Errorf("failed to get access token: failed to login {\"error\": \"repository error\"}"),
			IRepositoryMock: func(mc *minimock.Controller) repository.IRepository {
				mock := mocks.NewIRepositoryMock(mc)
				mock.GetAccessTokenMock.Expect(ctx, userID).Return(nil, repositoryError)
				return mock
			},
			AuthHelperMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				mock := mocks.NewAuthHelperMock(mc)
				mock.VerifyTokenMock.Expect(refreshToken).Return(&model.UserClaims{
					ID:   userID,
					Role: "user",
				}, nil)
				return mock
			},
			IProducerMockFunc: func(mc *minimock.Controller) kafka.IProducer {
				mock := mocks.NewIProducerMock(mc)
				return mock
			},
		},
		{
			name: "verify token error case",
			args: args{
				ctx: ctx,
				req: refreshToken,
			},
			want: nil,
			err:  fmt.Errorf("failed to verify token: failed to verify token {\"error\": \"verify token error\"}"),
			IRepositoryMock: func(mc *minimock.Controller) repository.IRepository {
				mock := mocks.NewIRepositoryMock(mc)
				return mock
			},
			AuthHelperMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				mock := mocks.NewAuthHelperMock(mc)
				mock.VerifyTokenMock.Expect(refreshToken).Return(&model.UserClaims{
					ID:   userID,
					Role: "user",
				}, verifyTokenError)
				return mock
			},
			IProducerMockFunc: func(mc *minimock.Controller) kafka.IProducer {
				mock := mocks.NewIProducerMock(mc)
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
			IProducerMock := tt.IProducerMockFunc(mc)

			myService := service.New(IRepoMock, AuthHelperMock, IProducerMock)

			result, err := myService.GetAccessToken(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
