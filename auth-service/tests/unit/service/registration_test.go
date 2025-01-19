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
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/repository/events"
	serviceAuth "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service/auth"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/db"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/tests/unit/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestRegistration(t *testing.T) {
	t.Parallel()
	logger.Init(logger.GetCore(logger.GetAtomicLevel(app.LogLevel)))

	type IRepositoryMockFunc func(mc *minimock.Controller) auth.IRepositoryAuth
	type TxManagerMockFunc func(mc *minimock.Controller) db.TxManager
	type IEventMockFunc func(mc *minimock.Controller) events.IEventRepository
	type JWTHelperMockFunc func(mc *minimock.Controller) jwt_manager.AuthHelper

	type args struct {
		ctx context.Context
		req *model.UserInfo
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, false, 8)
		name     = gofakeit.Name()
		surname  = gofakeit.LastName()
		role     = "user"

		accessToken  = gofakeit.UUID()
		refreshToken = gofakeit.UUID()

		req = &model.UserInfo{
			Email:    email,
			Password: password,
			Name:     name,
			Surname:  surname,
			Role:     role,
		}

		res = &model.AuthResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			UserID:       int64(0),
		}

		someError = fmt.Errorf("some error")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name          string
		args          args
		want          *model.AuthResponse
		err           error
		txMock        TxManagerMockFunc
		authRepoMock  IRepositoryMockFunc
		eventRepoMock IEventMockFunc
		jwtHelperMock JWTHelperMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			jwtHelperMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				mock := mocks.NewAuthHelperMock(mc)
				mock.HashPasswordMock.Expect(password).Return(password, nil)
				mock.GenerateAccessTokenMock.Expect(&model.AccessTokenInfo{
					ID:   0,
					Role: role,
				}).Return(accessToken, nil)
				mock.GenerateRefreshTokenMock.Expect(0).Return(refreshToken, nil)
				return mock
			},
			txMock: func(mc *minimock.Controller) db.TxManager {
				mock := mocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Return(nil)
				return mock
			},
			authRepoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				mock := mocks.NewIRepositoryAuthMock(mc)
				return mock
			},
			eventRepoMock: func(mc *minimock.Controller) events.IEventRepository {
				mock := mocks.NewIEventRepositoryMock(mc)
				return mock
			},
		},
		{
			name: "transaction failed",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  someError,
			jwtHelperMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				mock := mocks.NewAuthHelperMock(mc)
				mock.HashPasswordMock.Expect(password).Return(password, nil)
				return mock
			},
			txMock: func(mc *minimock.Controller) db.TxManager {
				mock := mocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Return(someError)
				return mock
			},
			authRepoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				mock := mocks.NewIRepositoryAuthMock(mc)
				return mock
			},
			eventRepoMock: func(mc *minimock.Controller) events.IEventRepository {
				mock := mocks.NewIEventRepositoryMock(mc)
				return mock
			},
		},
		{
			name: "generate access token failed",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  fmt.Errorf("failed to generate access token: %s", someError),
			jwtHelperMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				mock := mocks.NewAuthHelperMock(mc)
				mock.HashPasswordMock.Expect(password).Return(password, nil)
				mock.GenerateAccessTokenMock.Expect(&model.AccessTokenInfo{
					ID:   0,
					Role: role,
				}).Return(accessToken, someError)
				return mock
			},
			txMock: func(mc *minimock.Controller) db.TxManager {
				mock := mocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Return(nil)
				return mock
			},
			authRepoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				mock := mocks.NewIRepositoryAuthMock(mc)
				return mock
			},
			eventRepoMock: func(mc *minimock.Controller) events.IEventRepository {
				mock := mocks.NewIEventRepositoryMock(mc)
				return mock
			},
		},
		{
			name: "hash password failed",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  fmt.Errorf("failed to hash password: %s", someError),
			jwtHelperMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				mock := mocks.NewAuthHelperMock(mc)
				mock.HashPasswordMock.Expect(password).Return(password, someError)
				return mock
			},
			txMock: func(mc *minimock.Controller) db.TxManager {
				mock := mocks.NewTxManagerMock(mc)
				return mock
			},
			authRepoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				mock := mocks.NewIRepositoryAuthMock(mc)
				return mock
			},
			eventRepoMock: func(mc *minimock.Controller) events.IEventRepository {
				mock := mocks.NewIEventRepositoryMock(mc)
				return mock
			},
		},
		{
			name: "failed generate refresh token",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  fmt.Errorf("failed to generate refresh token: %s", someError),
			jwtHelperMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				mock := mocks.NewAuthHelperMock(mc)
				mock.HashPasswordMock.Expect(password).Return(password, nil)
				mock.GenerateAccessTokenMock.Expect(&model.AccessTokenInfo{
					ID:   0,
					Role: role,
				}).Return(accessToken, nil)
				mock.GenerateRefreshTokenMock.Expect(0).Return(refreshToken, someError)
				return mock
			},
			txMock: func(mc *minimock.Controller) db.TxManager {
				mock := mocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Return(nil)
				return mock
			},
			authRepoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				mock := mocks.NewIRepositoryAuthMock(mc)
				return mock
			},
			eventRepoMock: func(mc *minimock.Controller) events.IEventRepository {
				mock := mocks.NewIEventRepositoryMock(mc)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			IRepositoryMock := &repository.Repository{Auth: tt.authRepoMock(mc), Event: tt.eventRepoMock(mc)}
			service := serviceAuth.ServiceAuth{Repo: IRepositoryMock, AuthHelper: tt.jwtHelperMock(mc), Tx: tt.txMock(mc)}

			result, err := service.Registration(ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)

		})
	}

}
