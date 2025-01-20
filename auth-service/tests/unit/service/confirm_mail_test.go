package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/app"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/repository"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/repository/auth"
	serviceAuth "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service/auth"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/tests/unit/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestConfirmMail(t *testing.T) {
	t.Parallel()
	logger.Init(logger.GetCore(logger.GetAtomicLevel(app.LogLevel)))

	type IRepositoryMockFunc func(mc *minimock.Controller) auth.IRepositoryAuth

	type args struct {
		ctx  context.Context
		mail string
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		mail = gofakeit.Email()
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name         string
		args         args
		err          error
		authRepoMock IRepositoryMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx:  ctx,
				mail: mail,
			},
			err: nil,
			authRepoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				repo := mocks.NewIRepositoryAuthMock(mc)
				repo.ConfirmEmailMock.Expect(ctx, mail).Return(nil)
				return repo
			},
		},
		{
			name: "service error",
			args: args{
				ctx:  ctx,
				mail: mail,
			},
			err: fmt.Errorf("some error"),
			authRepoMock: func(mc *minimock.Controller) auth.IRepositoryAuth {
				repo := mocks.NewIRepositoryAuthMock(mc)
				repo.ConfirmEmailMock.Expect(ctx, mail).Return(fmt.Errorf("some error"))
				return repo
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			repositoryMock := &repository.Repository{Auth: tt.authRepoMock(mc)}
			service := serviceAuth.ServiceAuth{Repo: repositoryMock}

			err := service.ConfirmEmail(ctx, mail)
			require.Equal(t, tt.err, err)

		})
	}
}
