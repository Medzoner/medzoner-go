package services_test

import (
	"github.com/Medzoner/medzoner-go/pkg/infra/services"
	"github.com/sarulabs/di"
	"testing"
)

func TestGetDefinitions(t *testing.T) {
	t.Run("Unit: test services success", func(t *testing.T) {
		srvs := services.Service{}.GetDefinitions("../../../")
		builder, _ := di.NewBuilder()
		err := builder.Add(srvs...)
		if err != nil {
			panic(err)
		}
		_ = builder.Build()
	})
}
