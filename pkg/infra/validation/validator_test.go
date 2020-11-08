package validation_test

import (
	"github.com/Medzoner/medzoner-go/pkg/application/command"
	"github.com/Medzoner/medzoner-go/pkg/infra/validation"
	"gotest.tools/assert"
	"testing"
	"time"
)

func TestValidatetValidator(t *testing.T) {
	t.Run("Unit: test Validate success", func(t *testing.T) {
		validater := validation.ValidatorAdapter{}.New()

		createContactCommand := command.CreateContactCommand{
			DateAdd: time.Now(),
			Name:    "name",
			Email:   "email@email.email",
			Message: "message",
		}
		err := validater.Struct(createContactCommand)
		if err != nil {
			assert.Equal(t, true, false)
		}
		assert.Equal(t, true, true)
	})
	t.Run("Unit: test Validate failed", func(t *testing.T) {
		validater := validation.ValidatorAdapter{}.New()

		createContactCommand := command.CreateContactCommand{}
		err := validater.Struct(createContactCommand)
		if err == nil {
			assert.Equal(t, true, false)
		}
		assert.Equal(t, true, true)
		// @Todo
		validater.GetErrors()
		customErrors := validation.CustomError{}
		_ = customErrors.Error()
	})
}

func TestErrorValidator(t *testing.T) {
	t.Run("Unit: test Error success", func(t *testing.T) {
		validatorAdapter := validation.ValidatorAdapter{}.New()
		_ = validatorAdapter.GetErrors()
		assert.Equal(t, true, true)
	})
}
