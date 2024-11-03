package controller

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/pkg/auth_v1"
	"go.uber.org/zap"
)

// Registration handles user registration requests and returns a registration response.
func (c *Controller) Registration(ctx context.Context, req *authService.RegistrationRequest) (*authService.RegistrationResponse, error) {
	res, err := c.svc.Registration(ctx, converter.ToInfoFromProto(req))
	if err != nil {
		logger.Error(err.Error(), zap.Any("req", req))
		return nil, fmt.Errorf("failed to registration: %v", err)
	}

	return converter.ToProtoFromRegResponse(res), nil
}
