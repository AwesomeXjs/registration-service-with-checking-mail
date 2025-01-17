package service

//
//import (
//	"context"
//	"fmt"
//	"testing"
//
//	logger2 "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
//
//	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/kafka"
//	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/jwt_manager"
//	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
//	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/repository"
//	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service"
//	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/tests/unit/mocks"
//	"github.com/brianvoe/gofakeit"
//	"github.com/gojuno/minimock/v3"
//	"github.com/stretchr/testify/require"
//)
//
//const (
//	info = "info"
//)
//
//func TestConfirmEmail(t *testing.T) {
//	t.Parallel()
//	level := info
//	logger2.Init(logger2.GetCore(logger2.GetAtomicLevel(&level)))
//
//	type IRepositoryMockFunc func(mc *minimock.Controller) repository.IRepository
//	type AuthHelperMockFunc func(mc *minimock.Controller) jwt_manager.AuthHelper
//	type IProducerMockFunc func(mc *minimock.Controller) kafka.IProducer
//
//	type args struct {
//		ctx context.Context
//		req string
//	}
//
//	var (
//		ctx = context.Background()
//		mc  = minimock.NewController(t)
//
//		email = gofakeit.Email()
//
//		repositoryError = fmt.Errorf("repository error")
//	)
//
//	defer t.Cleanup(mc.Finish)
//
//	tests := []struct {
//		name              string              // название теста
//		args              args                // аргументы которые передаем в ручку Login
//		want              *model.AuthResponse // что хотим получить из ручки Login
//		err               error
//		IRepositoryMock   IRepositoryMockFunc // функция которая возвращает замоканый сервис с нужным поведением
//		AuthHelperMock    AuthHelperMockFunc
//		IProducerMockFunc IProducerMockFunc
//	}{
//		{
//			name: "success",
//			args: args{
//				ctx: ctx,
//				req: email,
//			},
//			want: nil,
//			err:  nil,
//			IRepositoryMock: func(mc *minimock.Controller) repository.IRepository {
//				mock := mocks.NewIRepositoryMock(mc)
//				// задаем поведение мока (все методы сервиса которые вызываются внутри ручки контроллера)
//				mock.ConfirmEmailMock.Expect(ctx, email).Return(nil)
//				return mock
//			},
//			AuthHelperMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
//				mock := mocks.NewAuthHelperMock(mc)
//				return mock
//			},
//			IProducerMockFunc: func(mc *minimock.Controller) kafka.IProducer {
//				mock := mocks.NewIProducerMock(mc)
//				return mock
//			},
//		},
//		{
//			name: "repository error case",
//			args: args{
//				ctx: ctx,
//				req: email,
//			},
//			want: nil,
//			err:  repositoryError,
//			IRepositoryMock: func(mc *minimock.Controller) repository.IRepository {
//				mock := mocks.NewIRepositoryMock(mc)
//				// задаем поведение мока (все методы сервиса которые вызываются внутри ручки контроллера)
//				mock.ConfirmEmailMock.Expect(ctx, email).Return(repositoryError)
//				return mock
//			},
//			AuthHelperMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
//				mock := mocks.NewAuthHelperMock(mc)
//				return mock
//			},
//			IProducerMockFunc: func(mc *minimock.Controller) kafka.IProducer {
//				mock := mocks.NewIProducerMock(mc)
//				return mock
//			},
//		},
//	}
//
//	for _, tt := range tests {
//		tt := tt
//		t.Run(tt.name, func(t *testing.T) {
//			t.Parallel()
//			IRepoMock := tt.IRepositoryMock(mc)
//			AuthHelperMock := tt.AuthHelperMock(mc)
//			IProducerMock := tt.IProducerMockFunc(mc)
//
//			myService := service.New(IRepoMock, AuthHelperMock, IProducerMock)
//
//			err := myService.ConfirmEmail(tt.args.ctx, tt.args.req)
//			require.Equal(t, tt.err, err)
//		})
//	}
//}
