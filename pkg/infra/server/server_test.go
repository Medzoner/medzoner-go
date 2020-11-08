package server_test

import (
	"context"
	"github.com/Medzoner/medzoner-go/pkg/infra/server"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	t.Run("Unit: test Server success", func(t *testing.T) {
		srv := server.Server{
			HTTPServer: &http.Server{
				Addr:    ":8124",
				Handler: &mux.Router{},
			},
		}
		go func() {
			_ = srv.ListenAndServe()
		}()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			cancel()
		}()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server Shutdown Failed:%+v", err)
		}
	})
	t.Run("Unit: test Server failed", func(t *testing.T) {
		srv := server.Server{}
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		_ = srv.ListenAndServe()
	})
}
