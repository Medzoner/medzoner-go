package validation

import "github.com/go-playground/validator/v10"

type MzValidator interface {
	New() MzValidator
	GetErrors() []CustomError
	Struct(str interface{}) error
}

type ValidatorAdapter struct {
	ValidationErrors []interface{}
	validatorLib     *validator.Validate
}

type CustomError struct {
	Tag string
}

func (v ValidatorAdapter) New() MzValidator {
	v.validatorLib = validator.New()
	return v
}

func (v ValidatorAdapter) GetErrors() []CustomError {
	var customErrors []CustomError
	var libErrors = validator.ValidationErrors{}
	for _, itemError := range libErrors {
		customErrors = append(customErrors, CustomError{Tag: itemError.ActualTag()})
	}
	return customErrors
}

func (v ValidatorAdapter) Struct(str interface{}) error {
	return v.validatorLib.Struct(str)
}

func (c *CustomError) Error() string {
	var resp = validator.ValidationErrors.Error(nil)
	return resp
}
