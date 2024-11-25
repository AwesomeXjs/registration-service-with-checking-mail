package service

import (
	"context"
)

func (s *Service) ConfirmEmail(ctx context.Context, mail string) error {
	return s.repo.ConfirmEmail(ctx, mail)
}
