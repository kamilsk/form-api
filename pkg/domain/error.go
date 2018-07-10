package domain

import "fmt"

// ValidationError defines the behavior of error related to invalid input values.
type ValidationError interface {
	error

	// HasError returns true if the input value has at least one error.
	HasError(input Input) bool
	// InputWithErrors returns a map of form inputs with at least one error and their errors.
	InputWithErrors() map[Input][]error
}

type validationError struct {
	single   bool
	position int
	value    string
	message  string
}

func (err validationError) Error() string {
	if err.single {
		if err.value != "" {
			return fmt.Sprintf("value %q at position %d is invalid: %s", err.value, err.position, err.message)
		}
		return fmt.Sprintf("value at position %d is invalid: %s", err.position, err.message)
	}
	return err.message
}

type dataValidationError struct {
	dataValidationResult
}

func (dataValidationError) Error() string {
	return "validation error"
}

func (err dataValidationError) HasError(input Input) bool {
	for _, result := range err.results {
		if result.input.Name == input.Name {
			return result.HasError()
		}
	}
	return false
}

func (err dataValidationError) InputWithErrors() map[Input][]error {
	m := make(map[Input][]error, len(err.results))
	for _, r := range err.results {
		if r.HasError() {
			m[r.input] = r.errors
		}
	}
	return m
}

type dataValidationResult struct {
	results []inputValidationResult
}

// AsError converts the result into error if it contains at least one input validation error.
func (r dataValidationResult) AsError() ValidationError {
	for _, result := range r.results {
		if result.HasError() {
			return dataValidationError{r}
		}
	}
	return nil
}

type inputValidationResult struct {
	input  Input
	errors []error
}

// HasError returns true if the result contains at least one error.
func (r inputValidationResult) HasError() bool {
	for _, err := range r.errors {
		if err != nil {
			return true
		}
	}
	return false
}
