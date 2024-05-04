package server_test

import (
	"context"
	"fmt"
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
			Logger:          &LoggerTest{},
			Router:          RouterMock{},
			IndexHandler:    nil,
			TechnoHandler:   nil,
			NotFoundHandler: nil,
			APIPort:         8123,
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

type LoggerTest struct {
	LogMessages []string
}

func (l *LoggerTest) Log(msg string) error {
	l.LogMessages = append(l.LogMessages, msg)
	fmt.Println(msg)
	return nil
}
func (l *LoggerTest) Error(msg string) error {
	l.LogMessages = append(l.LogMessages, msg)
	fmt.Println(msg)
	return nil
}

type RouterMock struct{}

func (r RouterMock) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	_ = path
	_ = f
	return &mux.Route{}
}

func (r RouterMock) PathPrefix(tpl string) *mux.Route {
	_ = tpl
	return &mux.Route{}
}

func (r RouterMock) Use(mwf ...mux.MiddlewareFunc) {
	_ = mwf
}

func (r RouterMock) SetNotFoundHandler(handler func(http.ResponseWriter, *http.Request)) {
	_ = handler
}

func (r RouterMock) ServeHTTP(http.ResponseWriter, *http.Request) {
}

func (r RouterMock) Handle(path string) {
	_ = path
}
