package interceptors

import (
	"context"
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()

	resp, err = handler(ctx, req)
	if err != nil {
		logger.Error(err.Error(), zap.String("method", info.FullMethod), zap.Any("req", req))
	}
	logger.Info("request", zap.String("method", info.FullMethod), zap.Any("req", req), zap.Any("resp", resp), zap.Duration("duration", time.Since(start)))
	return resp, err
}
