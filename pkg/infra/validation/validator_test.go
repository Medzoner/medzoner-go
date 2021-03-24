package validation_test

import (
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/infra/validation"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gotest.tools/assert"
	"reflect"
	"testing"
	"time"
)

func TestValidateValidator(t *testing.T) {
	t.Run("Unit: test Validate success", func(t *testing.T) {
		validatorAdapter := validation.ValidatorAdapter{}.New()

		createContactCommand := command.CreateContactCommand{
			DateAdd: time.Now(),
			Name:    "name",
			Email:   "email@email.email",
			Message: "message",
		}
		err := validatorAdapter.Struct(createContactCommand)
		if err != nil {
			assert.Equal(t, true, false)
		}
		assert.Equal(t, true, true)
	})
	t.Run("Unit: test Validate failed", func(t *testing.T) {
		validatorAdapter := validation.ValidatorAdapter{}.New()

		createContactCommand := command.CreateContactCommand{}
		err := validatorAdapter.Struct(createContactCommand)
		if err == nil {
			assert.Equal(t, true, false)
		}
		assert.Equal(t, true, true)
		// @Todo
		validatorAdapter.GetErrors()
		customErrors := validation.CustomError{}
		_ = customErrors.Error()
	})
}

func TestErrorValidator(t *testing.T) {
	t.Run("Unit: test Error success", func(t *testing.T) {
		var testFieldErrors = []validator.FieldError{FieldErrorTest{}}
		validatorAdapter := validation.ValidatorAdapter{
			ValidationErrors: testFieldErrors,
		}.New()
		_ = validatorAdapter.GetErrors()
		assert.Equal(t, true, true)
	})
}

func (f FieldErrorTest) Tag() string {
	return "tag1"
}

func (f FieldErrorTest) ActualTag() string {
	return "tag1"
}

func (f FieldErrorTest) Namespace() string {
	return "namespace1"
}

func (f FieldErrorTest) StructNamespace() string {
	return "namespace1"
}

func (f FieldErrorTest) Field() string {
	return "namespace1"
}

func (f FieldErrorTest) StructField() string {
	return "namespace1"
}

func (f FieldErrorTest) Value() interface{} {
	panic("implement me")
}

func (f FieldErrorTest) Param() string {
	return "namespace1"
}

func (f FieldErrorTest) Kind() reflect.Kind {
	panic("implement me")
}

func (f FieldErrorTest) Type() reflect.Type {
	panic("implement me")
}

func (f FieldErrorTest) Translate(ut ut.Translator) string {
	_ = ut
	return "ut"
}

func (f FieldErrorTest) Error() string {
	return "error"
}

type FieldErrorTest struct{}
