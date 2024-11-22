package kafka

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/logger"
	"go.uber.org/zap"
)

// Constants for environment variable names containing mail server configuration.
const (
	MailHost = "MAIL_HOST" // The hostname of the mail server.
	MailPort = "MAIL_PORT" // The port of the mail server.
	MailUser = "MAIL_USER" // The username for mail server authentication.
	MailPass = "MAIL_PASS" // The password for mail server authentication.
)

// IMailConfig defines the interface for accessing mail server configuration.
type IMailConfig interface {
	// GetHost returns the hostname of the mail server.
	GetHost() string
	// GetPort returns the port number of the mail server.
	GetPort() int
	// GetUserName returns the username used for mail server authentication.
	GetUserName() string
	// GetPassword returns the password used for mail server authentication.
	GetPassword() string
}

// MailConfig implements the IMailConfig interface and stores mail server configuration.
type MailConfig struct {
	host string // The hostname of the mail server.
	port int    // The port number of the mail server.
	user string // The username for mail server authentication.
	pass string // The password for mail server authentication.
}

// NewMailConfig creates a new instance of MailConfig by loading values from environment variables.
// Returns an error if any required environment variable is missing or invalid.
func NewMailConfig() (IMailConfig, error) {
	host := os.Getenv(MailHost)
	if len(host) == 0 {
		return nil, fmt.Errorf("failed to get mail host")
	}

	port, err := strconv.Atoi(os.Getenv(MailPort))
	if err != nil {
		return nil, fmt.Errorf("failed to get mail port")
	}

	user := os.Getenv(MailUser)
	if len(user) == 0 {
		return nil, fmt.Errorf("failed to get mail user")
	}

	pass := os.Getenv(MailPass)
	if len(pass) == 0 {
		logger.Error("failed to get mail pass", zap.String("mail pass", MailPass))
		return nil, fmt.Errorf("failed to get mail pass")
	}

	return &MailConfig{
		host: host,
		port: port,
		user: user,
		pass: pass,
	}, nil
}

// GetHost returns the hostname of the mail server.
func (m *MailConfig) GetHost() string {
	return m.host
}

// GetPort returns the port number of the mail server.
func (m *MailConfig) GetPort() int {
	return m.port
}

// GetUserName returns the username used for mail server authentication.
func (m *MailConfig) GetUserName() string {
	return m.user
}

// GetPassword returns the password used for mail server authentication.
func (m *MailConfig) GetPassword() string {
	return m.pass
}
