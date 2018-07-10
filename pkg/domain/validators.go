package domain

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Validate gets input metadata, validation rules and checks the input data.
func Validate(inputs []Input, rules map[string][]Validator, data map[string][]string) ValidationError {
	index := make(map[string]int, len(inputs))
	for i, input := range inputs {
		index[input.Name] = i
	}
	validation := dataValidationResult{}
	for name, values := range data {
		i, found := index[name]
		if !found {
			continue
		}
		inputValidation := inputValidationResult{input: inputs[i]}
		validators := rules[name]
		for _, validator := range validators {
			if err := validator.Validate(values); err != nil {
				inputValidation.errors = append(inputValidation.errors, err)
			}
		}
		validation.results = append(validation.results, inputValidation)
	}
	return validation.AsError()
}

// Validator defines the basic behavior of input validators.
type Validator interface {
	// Validate validates the input values.
	Validate(values []string) error
}

// ValidatorFunc type is an adapter to allow the use of ordinary functions as a validator.
type ValidatorFunc func(values []string) error

// Validate calls fn(values).
func (fn ValidatorFunc) Validate(values []string) error {
	return fn(values)
}

// LengthValidator returns the validator to check an input value length.
func LengthValidator(min, max int) ValidatorFunc {
	return func(values []string) error {
		for i, value := range values {
			if min != 0 && utf8.RuneCountInString(value) < min {
				return validationError{true, i, value, fmt.Sprintf("value length is less than %d", min)}
			}
			if max != 0 && utf8.RuneCountInString(value) > max {
				return validationError{true, i, value, fmt.Sprintf("value length is greater than %d", max)}
			}
		}
		return nil
	}
}

// RequireValidator returns the validator to check an input value for a not-empty.
func RequireValidator() ValidatorFunc {
	return func(values []string) error {
		if len(values) == 0 {
			return validationError{message: "values are empty"}
		}
		for i, value := range values {
			if value == "" {
				return validationError{single: true, position: i, message: "value is empty"}
			}
		}
		return nil
	}
}

// TypeValidator returns the validator to check an input value for compliance the type.
// It can raise the panic if the input type is unsupported.
func TypeValidator(inputType string, strict bool) ValidatorFunc {
	return func(values []string) error {
		switch inputType {
		case EmailType:
			for i, value := range values {
				// https://davidcel.is/posts/stop-validating-email-addresses-with-regex/
				if !strings.Contains(value, `@`) {
					return validationError{true, i, value, "value is not a valid email"}
				}
				if strict {
				}
			}
		case HiddenType, TextType:
			// nothing special
		default:
			panic(fmt.Sprintf("input type %q is not supported", inputType))
		}
		return nil
	}
}
