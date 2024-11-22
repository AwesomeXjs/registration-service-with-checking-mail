package grpc_server

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/mail_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *GrpcServer) CheckUniqCode(ctx context.Context, req *mail_v1.CheckUniqueCodeRequest) (*emptypb.Empty, error) {
	return nil, nil
}
