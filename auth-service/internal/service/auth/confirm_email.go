package auth

import (
	"context"
)

// ConfirmEmail is a service method that handles the email confirmation process.
// It calls the repository layer to update the email verification status.
func (s *ServiceAuth) ConfirmEmail(ctx context.Context, mail string) error {
	return s.repo.Auth.ConfirmEmail(ctx, mail)
}
