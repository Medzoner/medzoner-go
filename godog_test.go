package main

import (
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

	//quit := make(chan os.Signal)
	app := &pkg.App{}
	ctn := app.LoadContainer()

	go func() {
		log.Println("server starting")
		appWeb := ctn.Get("app-web").(*web.Web)
		appWeb.Start()
	}()
	fmt.Println("server started")

	baseUrl := "http://127.0.0.1:8000"
	status := godog.RunWithOptions("medzoner", func(s *godog.Suite) {
		bootstrap.New(baseUrl, app).FeatureContext(s)
	}, godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "pretty",
		Paths:  []string{"./features"},
		//Randomize: time.Now().UTC().UnixNano(),
	})

	if st := m.Run(); st > status {
		status = st
	}
	//
	//<-quit
	////// gracefully stop server
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancel()
	//if err := srv.Shutdown(ctx); err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("server stopped")
	os.Exit(status)
}

func TestRun(t *testing.T) {
	assert.Equal(t, 1, 1)
}
