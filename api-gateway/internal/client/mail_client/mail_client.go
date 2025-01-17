package mail_client

import (
	"context"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/api-gateway-auth/internal/model"
)

// MailClient defines the methods that a mail client must implement to interact
// with the mail service, specifically for checking the uniqueness of an email confirmation code.
type MailClient interface {
	// CheckUniqueCode checks whether the provided confirmation code is unique
	// for the given email and access token. It returns an error if the operation fails.
	CheckUniqueCode(ctx context.Context, accessToken string, request *model.ConfirmEmailRequest) error
}
