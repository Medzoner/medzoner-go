package definition_test

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/services/definition"
	"testing"
)

func TestConfigDefinition(t *testing.T) {
	t.Run("Unit: test config definition success", func(t *testing.T) {
		fmt.Println(definition.ConfigDependency{})
	})
}
