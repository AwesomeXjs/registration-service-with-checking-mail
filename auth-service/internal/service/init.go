package service

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/clients/kafka"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/repository"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/utils/auth_helper"
)

// Service provides business logic and interacts with the repository.
type Service struct {
	repo          repository.IRepository
	authHelper    auth_helper.AuthHelper
	kafkaProducer kafka.IProducer
}

// New creates a new Service instance with the given repository.
func New(repo repository.IRepository, authHelper auth_helper.AuthHelper, kafkaProducer kafka.IProducer) IService {
	return &Service{
		repo:          repo,
		authHelper:    authHelper,
		kafkaProducer: kafkaProducer,
	}
}
