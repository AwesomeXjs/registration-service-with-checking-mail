package service

import (
	"context"
	"testing"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/jwt_manager"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/repository"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/repository/auth"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/repository/events"
	serviceAuth "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service/auth"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/db"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/tests/unit/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestRegistration(t *testing.T) {
	t.Parallel()

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
		userID   = 1

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
			UserID:       int64(userID),
		}
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
				mock.GenerateAccessTokenMock.Expect(converter.ToModelAccessTokenInfo(userID, req)).Return(accessToken, nil)
				mock.GenerateRefreshTokenMock.Expect(userID).Return(refreshToken, nil)
				return mock
			},
			txMock: func(mc *minimock.Controller) db.TxManager {
				mock := mocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Expect(ctx, func(ctx context.Context) error { return nil }).Return(nil)
				return mock
			},
			authRepoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				mock := mocks.NewIRepositoryAuthMock(mc)
				mock.RegistrationMock.Expect(ctx, converter.FromUserInfoToDbModel(req, password)).Return(userID, nil)
				return mock
			},
			eventRepoMock: func(mc *minimock.Controller) events.IEventRepository {
				mock := mocks.NewIEventRepositoryMock(mc)
				mock.SendEventMock.Expect(ctx, converter.ToModelSendEvent("registration", []byte{1})).Return(nil)
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
