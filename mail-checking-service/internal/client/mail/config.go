package mail

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/pkg/logger"
	"go.uber.org/zap"
)

// Constants for environment variable names containing mail server configuration.
const (
	MyMailHost = "MAIL_HOST" // The hostname of the mail server.
	MyMailPort = "MAIL_PORT" // The port of the mail server.
	MyMailUser = "MAIL_USER" // The username for mail server authentication.
	MyMailPass = "MAIL_PASS" // The password for mail server authentication.
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

// MyMailConfig implements the IMailConfig interface and stores mail server configuration.
type MyMailConfig struct {
	host string // The hostname of the mail server.
	port int    // The port number of the mail server.
	user string // The username for mail server authentication.
	pass string // The password for mail server authentication.
}

// NewMailConfig creates a new instance of MailConfig by loading values from environment variables.
// Returns an error if any required environment variable is missing or invalid.
func NewMailConfig() (IMailConfig, error) {

	const mark = "Client.mail.NewMailConfig"

	host := os.Getenv(MyMailHost)
	if len(host) == 0 {
		logger.Error("failed to get mail host", mark, zap.String("mail host", MyMailHost))
		return nil, fmt.Errorf("failed to get mail host")
	}

	port, err := strconv.Atoi(os.Getenv(MyMailPort))
	if err != nil {
		logger.Error("failed to get mail port", mark, zap.String("mail port", MyMailPort))
		return nil, fmt.Errorf("failed to get mail port")
	}

	user := os.Getenv(MyMailUser)
	if len(user) == 0 {
		logger.Error("failed to get mail user", mark, zap.String("mail user", MyMailUser))
		return nil, fmt.Errorf("failed to get mail user")
	}

	pass := os.Getenv(MyMailPass)
	if len(pass) == 0 {
		logger.Error("failed to get mail pass", mark, zap.String("mail pass", MyMailPass))
		return nil, fmt.Errorf("failed to get mail pass")
	}

	return &MyMailConfig{
		host: host,
		port: port,
		user: user,
		pass: pass,
	}, nil
}

// GetHost returns the hostname of the mail server.
func (m *MyMailConfig) GetHost() string {
	return m.host
}

// GetPort returns the port number of the mail server.
func (m *MyMailConfig) GetPort() int {
	return m.port
}

// GetUserName returns the username used for mail server authentication.
func (m *MyMailConfig) GetUserName() string {
	return m.user
}

// GetPassword returns the password used for mail server authentication.
func (m *MyMailConfig) GetPassword() string {
	return m.pass
}
