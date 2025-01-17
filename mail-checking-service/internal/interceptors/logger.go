package interceptors

import (
	"context"
	"time"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// LogInterceptor is a gRPC interceptor that logs the details of incoming requests and their processing time.
func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	const mark = "Interceptors.LogInterceptor"

	start := time.Now()

	resp, err = handler(ctx, req)
	if err != nil {
		logger.Error(err.Error(), mark, zap.String("method", info.FullMethod), zap.Any("req", req))
	}
	logger.Info("request", mark, zap.String("method", info.FullMethod), zap.Any("req", req), zap.Any("resp", resp), zap.Duration("duration", time.Since(start)))
	return resp, err
}
