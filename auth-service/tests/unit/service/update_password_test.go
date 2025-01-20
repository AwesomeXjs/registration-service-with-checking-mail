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
	serviceAuth "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service/auth"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/tests/unit/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestUpdatePassword(t *testing.T) {
	t.Parallel()
	logger.Init(logger.GetCore(logger.GetAtomicLevel(app.LogLevel)))

	type IRepositoryMockFunc func(mc *minimock.Controller) auth.IRepositoryAuth
	type JWTHelperMockFunc func(mc *minimock.Controller) jwt_manager.AuthHelper

	type args struct {
		ctx context.Context
		req *model.UpdatePassInfo
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		email          = gofakeit.Email()
		password       = gofakeit.Password(true, true, true, true, false, 8)
		hashedPassword = gofakeit.Password(true, true, true, true, false, 8)

		req = &model.UpdatePassInfo{
			Email:       email,
			NewPassword: password,
		}
		someError = fmt.Errorf("some error")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name     string
		args     args
		err      error
		repoMock IRepositoryMockFunc
		jwtMock  JWTHelperMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: nil,
			repoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				repo := mocks.NewIRepositoryAuthMock(mc)
				repo.UpdatePasswordMock.Return(nil)
				return repo
			},
			jwtMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				jwt := mocks.NewAuthHelperMock(mc)
				jwt.HashPasswordMock.Expect(password).Return(hashedPassword, nil)
				return jwt
			},
		},
		{
			name: "hashed password error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: fmt.Errorf("failed to hash password: %v", someError),
			repoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				repo := mocks.NewIRepositoryAuthMock(mc)
				return repo
			},
			jwtMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				jwt := mocks.NewAuthHelperMock(mc)
				jwt.HashPasswordMock.Expect(password).Return(hashedPassword, someError)
				return jwt
			},
		},
		{
			name: "update password error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: fmt.Errorf("failed to update password: %v", someError),
			repoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				repo := mocks.NewIRepositoryAuthMock(mc)
				repo.UpdatePasswordMock.Return(someError)
				return repo
			},
			jwtMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				jwt := mocks.NewAuthHelperMock(mc)
				jwt.HashPasswordMock.Expect(password).Return(hashedPassword, nil)
				return jwt
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			IRepositoryMock := &repository.Repository{Auth: tt.repoMock(mc)}
			service := serviceAuth.ServiceAuth{Repo: IRepositoryMock, AuthHelper: tt.jwtMock(mc)}
			err := service.UpdatePassword(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
		})
	}
}
