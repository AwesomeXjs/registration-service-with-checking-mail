package service

import (
	"context"
	"strconv"
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

func TestRegistration(t *testing.T) {
	t.Parallel()
	level := "info"
	logger.Init(logger.GetCore(logger.GetAtomicLevel(&level)))

	type IRepositoryMockFunc func(mc *minimock.Controller) repository.IRepository
	type AuthHelperMockFunc func(mc *minimock.Controller) jwt_manager.AuthHelper
	type IProducerMockFunc func(mc *minimock.Controller) kafka.IProducer

	type args struct {
		ctx context.Context
		req *model.UserInfo
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		email        = gofakeit.Email()
		password     = gofakeit.Password(true, true, true, true, false, 8)
		name         = gofakeit.Name()
		surname      = gofakeit.LastName()
		role         = "user"
		hashPassword = gofakeit.UUID()

		accessToken  = "access-token"
		refreshToken = "refresh-token"
		userID       = 1

		topicRegistration = "registration"
		//timeNow           = time.Date(2025, 1, 16, 22, 49, 1, 593861534, time.Local)

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

		user = &model.InfoToDb{
			Email:        email,
			HashPassword: hashPassword,
			Role:         role,
		}

		//repositoryError = fmt.Errorf("repository error")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name              string              // название теста
		args              args                // аргументы которые передаем в ручку Login
		want              *model.AuthResponse // что хотим получить из ручки Login
		err               error
		IRepositoryMock   IRepositoryMockFunc // функция которая возвращает замоканый сервис с нужным поведением
		AuthHelperMock    AuthHelperMockFunc
		IProducerMockFunc IProducerMockFunc
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
				mock.RegistrationMock.Expect(ctx, user).Return(userID, nil)
				return mock
			},
			AuthHelperMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				mock := mocks.NewAuthHelperMock(mc)
				mock.HashPasswordMock.Expect(password).Return(hashPassword, nil)
				mock.GenerateAccessTokenMock.Expect(&model.AccessTokenInfo{
					ID:   userID,
					Role: role,
				}).Return(accessToken, nil)
				mock.GenerateRefreshTokenMock.Expect(userID).Return(refreshToken, nil)
				return mock
			},
			IProducerMockFunc: func(mc *minimock.Controller) kafka.IProducer {
				mock := mocks.NewIProducerMock(mc)
				mock.ProduceMock.Expect(email, topicRegistration, strconv.Itoa(userID)).Return(nil)
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

			result, err := myService.Registration(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
