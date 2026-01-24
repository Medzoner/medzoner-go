package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// MzValidator MzValidator
type MzValidator interface {
	// New() MzValidator
	GetErrors() []CustomError
	Struct(str interface{}) error
}

// ValidatorAdapter ValidatorAdapter
type ValidatorAdapter struct {
	validatorLib     *validator.Validate
	ValidationErrors []validator.FieldError
}

// NewValidatorAdapter NewValidatorAdapter
func NewValidatorAdapter() *ValidatorAdapter {
	v := ValidatorAdapter{
		ValidationErrors: validator.ValidationErrors{},
	}
	return v.New()
}

// CustomError CustomError
type CustomError struct {
	Tag string
}

// New New
func (v ValidatorAdapter) New() *ValidatorAdapter {
	v.validatorLib = validator.New()
	return &v
}

// GetErrors GetErrors
func (v ValidatorAdapter) GetErrors() []CustomError {
	customErrors := make([]CustomError, 0)
	for _, itemError := range v.ValidationErrors {
		customErrors = append(customErrors, CustomError{Tag: itemError.ActualTag()})
	}
	return customErrors
}

// Struct Struct
func (v ValidatorAdapter) Struct(str interface{}) error {
	if err := v.validatorLib.Struct(str); err != nil {
		return fmt.Errorf("error validating struct: %w", err)
	}
	return nil
}

// Error Error
func (c *CustomError) Error() string {
	return validator.ValidationErrors.Error(nil)
}
