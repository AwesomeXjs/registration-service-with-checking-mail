package mail

import (
	"context"
	"fmt"

	"gopkg.in/gomail.v2"
)

// IMailClient defines an interface for a mail client.
// Any structure implementing this interface must provide the SendEmail method.
type IMailClient interface {
	// SendEmail sends an email.
	// - ctx: Context for managing the lifecycle and cancellation of the operation.
	// - to: Recipient's email address.
	// - subject: Subject of the email.
	// - body: Content of the email.
	// Returns an error if the email fails to send.
	SendEmail(_ context.Context, to, subject, body string) error
}

// MyMailClient is a concrete implementation of the IMailClient interface.
// It uses configuration provided via the IMailConfig interface to send emails.
type MyMailClient struct {
	config IMailConfig // Configuration for connecting to the mail server.
}

// NewMailClient creates a new instance of MyMailClient.
// - config: An implementation of IMailConfig providing mail server details.
// Returns an IMailClient instance.
func NewMailClient(config IMailConfig) IMailClient {
	return &MyMailClient{
		config: config,
	}
}

// SendEmail sends an email using the provided configuration and details.
// - ctx: Context for managing lifecycle and cancellation.
// - config: Configuration for connecting to the mail server.
// - to: Recipient email address.
// - subject: Email subject.
// - body: Email body content.
// Returns an error if email sending fails.
func (m *MyMailClient) SendEmail(_ context.Context, to, subject, body string) error {
	// Create a new email message.
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "xxx@example.com")
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/plain", body)

	if m.config == nil {
		return fmt.Errorf("mail config is nil")
	}

	// Configure the email dialer with the provided server details.
	dialer := gomail.NewDialer(m.config.GetHost(), m.config.GetPort(), m.config.GetUserName(), m.config.GetPassword())

	// Attempt to send the email.
	if err := dialer.DialAndSend(mailer); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
