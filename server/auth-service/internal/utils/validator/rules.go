package validator

import (
	"context"
	"regexp"
	"unicode/utf8"
)

func ValidateName(name string) Condition {
	return func(ctx context.Context) error {
		if name == "" {
			return NewValidationErrors("empty name")
		}
		if utf8.RuneCountInString(name) < 3 || utf8.RuneCountInString(name) > 64 {
			return NewValidationErrors("name length must be between 3 and 64 characters")
		}
		return nil
	}
}

func ValidateRole(role string) Condition {
	return func(ctx context.Context) error {
		if role != "admin" && role != "user" {
			return NewValidationErrors("Only 'admin' and 'user' roles are allowed")
		}
		return nil
	}
}

func ValidateEmail(email string) Condition {
	return func(ctx context.Context) error {
		const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		re := regexp.MustCompile(emailRegex)
		result := re.MatchString(email)
		if !result {
			return NewValidationErrors("invalid email")
		}
		return nil
	}
}

func ValidatePassword(password string) Condition {
	return func(ctx context.Context) error {
		if utf8.RuneCountInString(password) < 8 || utf8.RuneCountInString(password) > 20 {
			return NewValidationErrors("password length must be between 8 and 20 characters")
		}
		return nil
	}
}

func ValidateSurname(surname string) Condition {
	return func(ctx context.Context) error {
		if surname == "" {
			return NewValidationErrors("empty surname")
		}
		if utf8.RuneCountInString(surname) < 3 || utf8.RuneCountInString(surname) > 64 {
			return NewValidationErrors("surname length must be between 3 and 64 characters")
		}
		return nil
	}
}
