package web_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/logger"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/web"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	t.Run("Unit: test Start failed on ListenAndServe", func(t *testing.T) {
		webTest := web.Web{
			Logger:          &LoggerTest{},
			Router:          RouterMock{},
			Server:          ServerFailedMock{},
			IndexHandler:    nil,
			TechnoHandler:   nil,
			ContactHandler:  nil,
			NotFoundHandler: nil,
			APIPort:         8123,
		}
		webTest.Start()
	})
	t.Run("Unit: test Start success", func(t *testing.T) {
		webTest := web.Web{
			Logger: &LoggerTest{},
			Router: RouterMock{},
			Server: &http.Server{
				Addr:    ":8123",
				Handler: RouterMock{},
			},
			IndexHandler:    nil,
			TechnoHandler:   nil,
			ContactHandler:  nil,
			NotFoundHandler: nil,
			APIPort:         8123,
		}
		go func() {
			webTest.Start()
		}()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			cancel()
		}()

		if err := webTest.Server.Shutdown(ctx); err != nil {
			log.Fatalf("Server Shutdown Failed:%+v", err)
		}
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
func (l LoggerTest) New() logger.ILogger {
	return &LoggerTest{}
}

type ServerFailedMock struct {
	HTTPServer *http.Server
}

func (s ServerFailedMock) ListenAndServe() error {
	return errors.New("failed ListenAndServe")
}

func (s ServerFailedMock) Shutdown(ctx context.Context) error {
	return errors.New("failed Shutdown")
}

type RouterMock struct{}

func (r RouterMock) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return &mux.Route{}
}

func (r RouterMock) PathPrefix(tpl string) *mux.Route {
	return &mux.Route{}
}

func (r RouterMock) Use(mwf ...mux.MiddlewareFunc) {
}

func (r RouterMock) SetNotFoundHandler(handler func(http.ResponseWriter, *http.Request)) {
}

func (r RouterMock) ServeHTTP(http.ResponseWriter, *http.Request) {
}

func (r RouterMock) Handle(path string) {
}
