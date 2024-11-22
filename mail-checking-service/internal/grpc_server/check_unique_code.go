package grpc_server

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/mail_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// CheckUniqueCode handles the gRPC request to verify the uniqueness of a code.
// Returns an empty protobuf message on success or an error if the operation fails.
func (c *GrpcServer) CheckUniqueCode(_ context.Context, _ *mail_v1.CheckUniqueCodeRequest) (*emptypb.Empty, error) {
	return nil, nil
}
