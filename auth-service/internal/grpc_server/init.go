package grpc_server

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service"
	authService "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/auth_v1"
)

// GrpcServer handles gRPC requests for the authentication service.
// It embeds UnimplementedAuthV1Server to provide default implementations
// for the gRPC server methods, allowing customization where needed.
type GrpcServer struct {
	authService.UnimplementedAuthV1Server
	svc service.IService
}

// New creates a new instance of Controller with the provided service.
// The service is used to manage authentication-related operations.
func New(svc service.IService) *GrpcServer {
	return &GrpcServer{
		svc: svc,
	}
}
