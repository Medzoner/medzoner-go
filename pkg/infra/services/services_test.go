package services_test

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/services"
	"testing"
)

func TestApp(t *testing.T) {
	t.Run("Unit: test services success", func(t *testing.T) {
		services.Service{}.GetDefinitions("../../../")
	})
}
