package grpc_server

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/logger"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ConfirmEmail is a gRPC handler that processes email confirmation requests.
// It calls the service layer to mark the email as verified and returns an error if the operation fails.
func (c *GrpcServer) ConfirmEmail(ctx context.Context, req *authService.ConfirmEmailRequest) (*emptypb.Empty, error) {
	err := c.svc.ConfirmEmail(ctx, req.GetEmail())
	if err != nil {
		logger.Error("failed to confirm email", zap.Error(err))
		return nil, fmt.Errorf("failed to confirm email: %v", err)
	}

	return nil, nil
}
