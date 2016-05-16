package definition_test

import (
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/services/definition"
	"testing"
)

func TestDefLogger(t *testing.T) {
	t.Run("Unit: test service log success", func(t *testing.T) {
		fmt.Println(definition.LoggerDefinition)
	})
}
