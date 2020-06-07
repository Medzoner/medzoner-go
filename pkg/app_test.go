package pkg_test

import (
	"context"
	"github.com/Medzoner/medzoner-go/pkg"
	"testing"
	"time"
)

func TestHandle(t *testing.T) {
	t.Run("Unit: test App success", func(t *testing.T) {
		app := pkg.App{
			DebugMode: false,
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
}
