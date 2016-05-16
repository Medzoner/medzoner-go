package definition_test

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/services/definition"
	"testing"
)

func TestDefConfig(t *testing.T) {
	t.Run("Unit: test service config success", func(t *testing.T) {
		fmt.Println(definition.ConfigDependency{})
	})
}
