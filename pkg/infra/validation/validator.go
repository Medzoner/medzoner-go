package validation

import (
	"github.com/go-playground/validator/v10"
)

//MzValidator MzValidator
type MzValidator interface {
	New() MzValidator
	GetErrors() []CustomError
	Struct(str interface{}) error
}

//ValidatorAdapter ValidatorAdapter
type ValidatorAdapter struct {
	ValidationErrors []validator.FieldError
	validatorLib     *validator.Validate
}

//CustomError CustomError
type CustomError struct {
	Tag string
}

//New New
func (v ValidatorAdapter) New() MzValidator {
	v.validatorLib = validator.New()
	return v
}

//GetErrors GetErrors
func (v ValidatorAdapter) GetErrors() []CustomError {
	var customErrors []CustomError
	for _, itemError := range v.ValidationErrors {
		customErrors = append(customErrors, CustomError{Tag: itemError.ActualTag()})
	}
	return customErrors
}

//Struct Struct
func (v ValidatorAdapter) Struct(str interface{}) error {
	return v.validatorLib.Struct(str)
}

//Error Error
func (c *CustomError) Error() string {
	var resp = validator.ValidationErrors.Error(nil)
	return resp
}
