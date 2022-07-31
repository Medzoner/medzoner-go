package definition_test

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/services/definition"
	"testing"
)

func TestHandlerDefinition(t *testing.T) {
	t.Run("Unit: test config definition success", func(t *testing.T) {
		fmt.Println(definition.NotFoundHandlerDefinition)
		fmt.Println(definition.IndexHandlerDefinition)
		fmt.Println(definition.TechnoHandlerDefinition)
	})
}
