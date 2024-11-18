package validator

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// ValidationErrors holds validation error messages.
type ValidationErrors struct {
	Messages []string `json:"error_messages"`
}

// addError appends a new error message to the list of validation errors.
func (v *ValidationErrors) addError(message string) {
	v.Messages = append(v.Messages, message)
}

// NewValidationErrors creates a new instance of ValidationErrors with given messages.
func NewValidationErrors(messages ...string) *ValidationErrors {
	return &ValidationErrors{
		Messages: messages,
	}
}

// Error returns a JSON string representation of the validation error messages.
// If marshaling fails, it returns the marshaling error as a string.
func (v *ValidationErrors) Error() string {
	data, err := json.Marshal(v.Messages)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

// IsValidationError checks if the provided error is of type ValidationErrors.
// It returns true if the error is a validation error.
func IsValidationError(err error) bool {
	var ve *ValidationErrors
	return errors.As(err, &ve)
}
