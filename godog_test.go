package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Medzoner/medzoner-go/features/bootstrap"
	"github.com/Medzoner/medzoner-go/pkg/infra/dependency"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"gotest.tools/assert"
	"log"
	"os"
	"sync"
	"testing"
	"time"
)

var opt = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress",
}

func init() {
	godog.BindCommandLineFlags("godog.", &opt)
}

func TestMain(m *testing.M) {
	flag.Parse()

	srv := dependency.InitServer()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		log.Println("server starting")
		wg.Done()
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	wg.Wait()
	fmt.Println("server started")

	baseURL := "http://127.0.0.1:8002"

	opts := godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "pretty",
		Paths:  []string{"./features"},
		//Randomize: time.Now().UTC().UnixNano(),
	}

	featureCtx := bootstrap.New(baseURL)
	status := godog.TestSuite{
		Name: "medzoner",
		TestSuiteInitializer: func(suiteContext *godog.TestSuiteContext) {
			featureCtx.InitializeTestSuite(suiteContext)
		},
		ScenarioInitializer: func(scenarioContext *godog.ScenarioContext) {
			featureCtx.InitializeScenario(scenarioContext)
		},
		Options: &opts,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	go func() {
		log.Println("server stopping")
		if err := srv.Shutdown(ctx); err != nil {
			log.Println(err)
		}
	}()
	log.Println("server stopped with status: ", status)
	os.Exit(status)
}

func TestRun(t *testing.T) {
	assert.Equal(t, 1, 1)
}
