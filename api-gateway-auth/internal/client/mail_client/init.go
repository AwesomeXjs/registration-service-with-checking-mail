package mail_client

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/mail_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCMailClient struct {
	mailClient MailClient
}

func NewGRPCMailClient(mailClient MailClient) MailClient {
	return &GRPCMailClient{
		mailClient: mailClient,
	}
}

func (g *GRPCMailClient) CheckUniqueCode(ctx context.Context, in *mail_v1.CheckUniqueCodeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.mailClient.CheckUniqueCode(ctx, in, opts...)
}
