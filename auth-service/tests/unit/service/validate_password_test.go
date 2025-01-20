package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/app"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/jwt_manager"
	serviceAuth "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service/auth"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/tests/unit/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestValidatePassword(t *testing.T) {
	t.Parallel()
	logger.Init(logger.GetCore(logger.GetAtomicLevel(app.LogLevel)))

	type JWTHelperMockFunc func(mc *minimock.Controller) jwt_manager.AuthHelper

	type args struct {
		ctx         context.Context
		accessToken string
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		accessToken = gofakeit.UUID()

		someError = fmt.Errorf("some error")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name    string
		args    args
		want    bool
		err     error
		jwtMock JWTHelperMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx:         ctx,
				accessToken: accessToken,
			},
			want: true,
			err:  nil,
			jwtMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				jwtMock := mocks.NewAuthHelperMock(mc)
				jwtMock.VerifyTokenMock.Expect(accessToken).Return(nil, nil)
				return jwtMock
			},
		},
		{
			name: "verify token error",
			args: args{
				ctx:         ctx,
				accessToken: accessToken,
			},
			want: false,
			err:  fmt.Errorf("failed to verify token: %v", someError),
			jwtMock: func(mc *minimock.Controller) jwt_manager.AuthHelper {
				jwtMock := mocks.NewAuthHelperMock(mc)
				jwtMock.VerifyTokenMock.Expect(accessToken).Return(nil, someError)
				return jwtMock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			service := serviceAuth.ServiceAuth{AuthHelper: tt.jwtMock(mc)}
			result, err := service.ValidateToken(ctx, tt.args.accessToken)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
