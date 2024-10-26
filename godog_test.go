package main

import (
	"context"
	"os"
	"testing"

	"github.com/Medzoner/medzoner-go/features/bootstrap"
	"github.com/Medzoner/medzoner-go/pkg/infra/dependency"
	mocks "github.com/Medzoner/medzoner-go/test"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/golang/mock/gomock"
	metricNoop "go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/trace/noop"
)

var opt = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress",
}

func init() {
	godog.BindCommandLineFlags("godog.", &opt)
}

func TestFeatures(t *testing.T) {
	mocked := mocks.New(t)
	mocked.ContactRepository.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	mocked.HttpTracer.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
	mocked.HttpTracer.EXPECT().Int64Counter(gomock.Any(), gomock.Any()).Return(metricNoop.Int64Counter{}, nil).AnyTimes()
	mocked.HttpTracer.EXPECT().WriteLog(gomock.Any(), gomock.Any()).Return().AnyTimes()
	mocked.Mailer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()
	mocked.TechnoRepository.EXPECT().FetchStack().Return(map[string]interface{}{}, nil).AnyTimes()

	_ = os.Setenv("APP_ENV", "test")
	_ = os.Setenv("DEBUG", "true")
	srv, err := dependency.InitServerTest(&mocked)
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
	featureCtx := bootstrap.New(*srv, mocked)
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
