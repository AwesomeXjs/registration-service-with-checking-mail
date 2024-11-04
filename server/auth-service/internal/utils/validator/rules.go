package validator

import (
	"context"
	"regexp"
	"unicode/utf8"
)

// ValidateName checks the validity of a name string.
// It ensures that the name is not empty and has a length between 3 and 64 characters.
func ValidateName(name string) Condition {
	return func(_ context.Context) error {
		if name == "" {
			return NewValidationErrors("empty name")
		}
		if utf8.RuneCountInString(name) < 3 || utf8.RuneCountInString(name) > 64 {
			return NewValidationErrors("name length must be between 3 and 64 characters")
		}
		return nil
	}
}

// ValidateRole checks if the provided role is valid.
// It only allows "admin" or "user" as acceptable roles.
func ValidateRole(role string) Condition {
	return func(_ context.Context) error {
		if role != "admin" && role != "user" {
			return NewValidationErrors("Only 'admin' and 'user' roles are allowed")
		}
		return nil
	}
}

// ValidateEmail checks if the given email string is in a valid format.
// It uses a regular expression to ensure the email is correctly structured.
func ValidateEmail(email string) Condition {
	return func(_ context.Context) error {
		const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		re := regexp.MustCompile(emailRegex)
		result := re.MatchString(email)
		if !result {
			return NewValidationErrors("invalid email")
		}
		return nil
	}
}

// ValidatePassword checks if the provided password meets the length requirement.
// It ensures the password length is between 8 and 20 characters.
func ValidatePassword(password string) Condition {
	return func(_ context.Context) error {
		if utf8.RuneCountInString(password) < 8 || utf8.RuneCountInString(password) > 20 {
			return NewValidationErrors("password length must be between 8 and 20 characters")
		}
		return nil
	}
}

// ValidateSurname checks the validity of a surname string.
// It ensures that the surname is not empty and has a length between 3 and 64 characters.
func ValidateSurname(surname string) Condition {
	return func(_ context.Context) error {
		if surname == "" {
			return NewValidationErrors("empty surname")
		}
		if utf8.RuneCountInString(surname) < 3 || utf8.RuneCountInString(surname) > 64 {
			return NewValidationErrors("surname length must be between 3 and 64 characters")
		}
		return nil
	}
}
