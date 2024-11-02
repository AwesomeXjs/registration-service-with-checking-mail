package service

import "github.com/AwesomeXjs/registration-service-with-checking-mail/server/auth-service/internal/repository"

type Service struct {
	repo repository.IRepository
}

func New(repo repository.IRepository) *Service {
	return &Service{
		repo: repo,
	}
}
