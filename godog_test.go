package main

import (
	"context"
	"os"
	"testing"

	"github.com/Medzoner/medzoner-go/features/bootstrap"
	mocks "github.com/Medzoner/medzoner-go/test"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/golang/mock/gomock"
	"github.com/Medzoner/medzoner-go/internal/wire"
)

var opt = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress",
}

func init() {
	godog.BindCommandLineFlags("godog.", &opt)
}

func TestFeatures(t *testing.T) {
	disable := true

	if disable {
		t.Skip("Skipping godog tests")
		return
	}

	mocked := mocks.New(t)
	mocked.ContactRepository.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	mocked.Mailer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()
	mocked.TechnoRepository.EXPECT().FetchStack(context.Background()).Return(map[string]interface{}{}, nil).AnyTimes()

	t.Setenv("APP_ENV", "test")
	t.Setenv("DEBUG", "true")
	srv, err := wire.InitServerTest(context.Background(), mocked)
	if err != nil {
		t.Error(err)
		return
	}

	opts := godog.Options{
		Output:      colors.Colored(os.Stdout),
		Format:      "pretty",
		Paths:       []string{"./features"},
		Concurrency: 4,
	}
	featureCtx := bootstrap.New(srv, *mocked)
	suite := godog.TestSuite{
		Name: "medzoner",
		TestSuiteInitializer: func(suiteContext *godog.TestSuiteContext) {
			featureCtx.InitializeTestSuite(suiteContext)
		},
		ScenarioInitializer: func(scenarioContext *godog.ScenarioContext) {
			featureCtx.InitializeScenario(scenarioContext)
		},
		Options: &opts,
	}

	status := suite.Run()

	if status != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
