package web_test

import (
	"context"
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
	t.Run("Unit: test Start success", func(t *testing.T) {
		router := &mux.Router{}
		webTest := web.Web{
			Logger:         &LoggerTest{},
			Router:         router,
			Server:         &http.Server{
				Addr:    ":8123",
				Handler: router,
			},
			IndexHandler:   nil,
			TechnoHandler:  nil,
			ContactHandler: nil,
			ApiPort:        8123,
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
func (l *LoggerTest) Log(msg string) {
	l.LogMessages = append(l.LogMessages, msg)
	fmt.Println(msg)
}
func (l *LoggerTest) Error(msg string) {
	l.LogMessages = append(l.LogMessages, msg)
	fmt.Println(msg)
}
func (l LoggerTest) New() logger.ILogger {
	return &LoggerTest{}
}