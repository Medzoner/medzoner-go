package app_test

import (
	"context"
	wiring "github.com/Medzoner/medzoner-go/pkg/infra/dependency"
	"testing"
	"time"
)

func TestHandle(t *testing.T) {
	t.Run("Unit: test App success migrate up", func(t *testing.T) {
		appli := wiring.InitApp()
		//builder, _ := di.NewBuilder()
		//appli.LoadContainer(builder)
		// @Todo Mock db
		appli.Handle("migrate-up")
	})
	t.Run("Unit: test App failed", func(t *testing.T) {
		appli := wiring.InitApp()
		//builder, _ := di.NewBuilder()
		//appli.LoadContainer(builder)
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		appli.Handle("web")
	})
	t.Run("Unit: test App success web server", func(t *testing.T) {
		appli := wiring.InitApp()
		//builder, _ := di.NewBuilder()
		//appli.LoadContainer(builder)
		go func() {
			appli.Handle("web")
		}()
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			cancel()
		}()
	})
	t.Run("Unit: test App success unknown", func(t *testing.T) {
		appli := wiring.InitApp()
		//builder, _ := di.NewBuilder()
		//appli.LoadContainer(builder)
		go func() {
			appli.Handle("unknown")
		}()
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			cancel()
		}()
	})
	t.Run("Unit: test App success migrate", func(t *testing.T) {
		appli := wiring.InitApp()
		//builder, _ := di.NewBuilder()
		//appli.LoadContainer(builder)
		//defer func() {
		//	if r := recover(); r == nil {
		//		t.Errorf("The code did not panic")
		//	}
		//}()
		appli.Handle("migrate")
	})
}
