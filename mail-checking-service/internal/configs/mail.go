package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/mail-checking-service/internal/utils/logger"
	"go.uber.org/zap"
)

const (
	MailHost = "MAIL_HOST"
	MailPort = "MAIL_PORT"
	MailUser = "MAIL_USER"
	MailPass = "MAIL_PASS"
)

type IMailConfig interface {
	GetHost() string
	GetPort() int
	GetUserName() string
	GetPassword() string
}

type MailConfig struct {
	host string
	port int
	user string
	pass string
}

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

func (m *MailConfig) GetHost() string {
	return m.host
}

func (m *MailConfig) GetPort() int {
	return m.port
}

func (m *MailConfig) GetUserName() string {
	return m.user
}

func (m *MailConfig) GetPassword() string {
	return m.pass
}
