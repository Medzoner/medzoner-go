package server_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/server"
	tracerMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/infra/tracer"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	t.Run("Unit: test Server success", func(t *testing.T) {
		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().ShutdownMeter(gomock.Any()).Return(nil).AnyTimes()
		httpTracerMock.EXPECT().ShutdownTracer(gomock.Any()).Return(nil).AnyTimes()
		httpTracerMock.EXPECT().ShutdownLogger(gomock.Any()).Return(nil).AnyTimes()
		srv := server.NewServer(config.Config{APIPort: 8123}, RouterMock{}, &LoggerTest{}, httpTracerMock)
		go func() {
			srv.Start(context.Background())
		}()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			t.Errorf("Server Shutdown Failed:%+v", err)
		}
	})
	t.Run("Unit: test Server error", func(t *testing.T) {
		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().ShutdownMeter(gomock.Any()).Return(errors.New("error")).AnyTimes()
		httpTracerMock.EXPECT().ShutdownTracer(gomock.Any()).Return(errors.New("error")).AnyTimes()
		httpTracerMock.EXPECT().ShutdownLogger(gomock.Any()).Return(errors.New("error")).AnyTimes()
		srv := server.NewServer(config.Config{APIPort: 8123}, RouterMock{}, &LoggerTest{}, httpTracerMock)
		go func() {
			srv.Start(context.Background())
		}()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			t.Errorf("Server Shutdown Failed:%+v", err)
		}
	})
}

type LoggerTest struct {
	LogMessages []string
}

func (l *LoggerTest) Log(msg string) {
	l.LogMessages = append(l.LogMessages, msg)
	fmt.Println(msg)
}
func (l *LoggerTest) Error(msg string) {
	l.LogMessages = append(l.LogMessages, msg)
	fmt.Println(msg)
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
