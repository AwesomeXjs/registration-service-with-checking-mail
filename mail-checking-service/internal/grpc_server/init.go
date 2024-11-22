package grpc_server

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/mail_v1"
)

// GrpcServer handles gRPC requests for the authentication service.
// It embeds UnimplementedMailV1Server to provide default implementations
// for the gRPC server methods, allowing customization where needed.
type GrpcServer struct {
	mail_v1.UnimplementedMailV1Server
}

// New creates a new instance of Controller with the provided service.
// The service is used to manage authentication-related operations.
func New() *GrpcServer {
	return &GrpcServer{}
}
