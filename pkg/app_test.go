package pkg_test

import (
	"context"
	"github.com/Medzoner/medzoner-go/pkg"
	"testing"
	"time"
)

func TestHandle(t *testing.T) {
	t.Run("Unit: test App success web server", func(t *testing.T) {
		app := pkg.App{
			DebugMode: true,
			RootPath:  "../",
		}
		go func() {
			app.Handle("web")
		}()
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			cancel()
		}()
	})
	t.Run("Unit: test App success migrate up", func(t *testing.T) {
		app := pkg.App{
			DebugMode: true,
			RootPath:  "../",
		}
		app.Handle("migrate-up")
	})
	t.Run("Unit: test App failed", func(t *testing.T) {
		app := pkg.App{
			DebugMode: true,
			RootPath:  "../fake",
		}
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		app.Handle("web")
	})
}
