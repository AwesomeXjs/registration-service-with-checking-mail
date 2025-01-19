package mail_client

import (
	"context"
	"fmt"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/converter"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/model"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/pkg/logger"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/mail_v1"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// GRPCMailClient is a struct that holds an instance of the mail_v1.MailV1Client.
// It provides gRPC-based communication methods for mail-related operations.
type GRPCMailClient struct {
	mailClient mail_v1.MailV1Client // The actual gRPC client that communicates with the mail service.
}

// NewGRPCMailClient creates and returns a new instance of GRPCMailClient.
// This function wraps the provided mail_v1.MailV1Client instance, allowing gRPC communication.
func NewGRPCMailClient(mailClient mail_v1.MailV1Client) MailClient {
	return &GRPCMailClient{
		mailClient: mailClient, // Assign the provided gRPC mail client to the GRPCMailClient.
	}
}

// CheckUniqueCode calls the CheckUniqueCode method on the gRPC mail client.
// It takes an access token and a ConfirmEmailRequest, converts the request to the appropriate proto format,
// and then checks if the unique code is valid. If the operation fails, it logs an error and returns it.
func (g *GRPCMailClient) CheckUniqueCode(ctx context.Context, accessToken string, request *model.ConfirmEmailRequest) error {

	const mark = "Client.mail_client.CheckUniqueCode"

	span, contextWithTrace := opentracing.StartSpanFromContext(ctx, "confirm email")
	defer span.Finish()

	span.SetTag("email", request.Email)

	// Convert the model request to the proto request format and call CheckUniqueCode on the gRPC client.
	_, err := g.mailClient.CheckUniqueCode(contextWithTrace, converter.FromModelToProtoCheckUniqueCode(request, accessToken))
	if err != nil {
		// Log the error if the gRPC call fails.
		logger.Error("failed to check unique code", mark, zap.Error(err))
		// Return a formatted error.
		return fmt.Errorf("failed to check unique code: %v", err)
	}
	// Return nil if no error occurred.
	return nil
}
