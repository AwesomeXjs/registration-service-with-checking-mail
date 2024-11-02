package service

import "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/repository"

// Service provides business logic and interacts with the repository.
type Service struct {
	repo repository.IRepository
}

// New creates a new Service instance with the given repository.
func New(repo repository.IRepository) *Service {
	return &Service{
		repo: repo,
	}
}
