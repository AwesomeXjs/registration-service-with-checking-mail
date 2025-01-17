package service

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/kafka"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/jwt_manager"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/repository"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service/auth"
	eventSender "github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/service/event_sender"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/pkg/db"
)

// Service provides business logic and interacts with the repository.
type Service struct {
	Auth  auth.IServiceAuth
	Event eventSender.ISender
}

// New creates a new Service instance with the given repository.
func New(repo *repository.Repository,
	authHelper jwt_manager.AuthHelper,
	tx db.TxManager, producer kafka.IProducer, db db.Client) *Service {
	return &Service{
		Auth: auth.New(repo, authHelper, tx),
		Event: eventSender.New(
			producer,
			db,
			repo.Event,
		),
	}
}
