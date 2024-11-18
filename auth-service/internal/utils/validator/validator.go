package validator

import "context"

// Condition defines a function type that takes a context and returns an error.
type Condition func(ctx context.Context) error

// Validate checks a series of conditions and collects validation errors.
// If any condition returns a validation error, it is added to the validation errors collection.
// If any other error occurs, it is returned immediately.
func Validate(ctx context.Context, conds ...Condition) error {
	ve := NewValidationErrors()

	for _, c := range conds {
		err := c(ctx)
		if err != nil {
			if IsValidationError(err) {
				ve.addError(err.Error())
				continue
			}

			return err
		}
	}

	if ve.Messages == nil {
		return nil
	}

	return ve
}
