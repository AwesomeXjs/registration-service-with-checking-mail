package service

import (
	"context"
)

// ConfirmEmail is a service method that handles the email confirmation process.
// It calls the repository layer to update the email verification status.
func (s *Service) ConfirmEmail(ctx context.Context, mail string) error {
	return s.repo.ConfirmEmail(ctx, mail)
}
