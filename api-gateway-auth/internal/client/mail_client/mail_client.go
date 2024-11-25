package mail_client

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/mail_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MailClient interface {
	CheckUniqueCode(ctx context.Context, in *mail_v1.CheckUniqueCodeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}
