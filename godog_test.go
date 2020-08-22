package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
	"github.com/Medzoner/medzoner-go/features/bootstrap"
	"github.com/Medzoner/medzoner-go/pkg"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/web"
	"gotest.tools/assert"
	"log"
	"os"
	"testing"
	"time"
)

var opt = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress",
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	flag.Parse()

	rootPath, _ := os.Getwd()
	app := &pkg.App{
		RootPath: rootPath,
	}
	ctn := app.LoadContainer()

	appWeb := ctn.Get("app-web").(*web.Web)
	go func() {
		log.Println("server starting")
		appWeb.Start()
	}()
	fmt.Println("server started")

	baseURL := "http://127.0.0.1:8000"
	status := godog.RunWithOptions("medzoner", func(s *godog.Suite) {
		bootstrap.New(baseURL, app).FeatureContext(s)
	}, godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "pretty",
		Paths:  []string{"./features"},
		//Randomize: time.Now().UTC().UnixNano(),
	})

	if st := m.Run(); st > status {
		status = st
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := appWeb.Server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("server stopped")
	os.Exit(status)
}

func TestRun(t *testing.T) {
	assert.Equal(t, 1, 1)
}
