package service

import (
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/repository"
	"github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/utils/auth_helper"
)

// Service provides business logic and interacts with the repository.
type Service struct {
	repo       repository.IRepository
	authHelper auth_helper.AuthHelper
}

// New creates a new Service instance with the given repository.
func New(repo repository.IRepository, authHelper auth_helper.AuthHelper) IService {
	return &Service{
		repo:       repo,
		authHelper: authHelper,
	}
}
