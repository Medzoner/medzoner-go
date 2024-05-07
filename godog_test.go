package main

import (
	"github.com/Medzoner/medzoner-go/features/bootstrap"
	"github.com/Medzoner/medzoner-go/pkg/infra/dependency"
	mocks "github.com/Medzoner/medzoner-go/test"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/golang/mock/gomock"
	"os"
	"testing"
)

var opt = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress",
}

func init() {
	godog.BindCommandLineFlags("godog.", &opt)
}

func TestFeatures(t *testing.T) {
	mockedRepository := mocks.New(t)
	srv := dependency.InitServerTest(mockedRepository)
	mockedRepository.ContactRepository.EXPECT().Save(gomock.Any(), gomock.Any()).Return()

	opts := godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "pretty",
		Paths:  []string{"./features"},
		//Randomize: time.Now().UTC().UnixNano(),
	}

	featureCtx := bootstrap.New(*srv)
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

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
