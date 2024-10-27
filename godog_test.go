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
	mocked.HttpTracer.EXPECT().StartRoot(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
	mocked.HttpTracer.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
	mocked.HttpTracer.EXPECT().ShutdownTracer(gomock.Any()).Return(nil).AnyTimes()
	mocked.HttpTracer.EXPECT().ShutdownMeter(gomock.Any()).Return(nil).AnyTimes()
	mocked.HttpTracer.EXPECT().ShutdownLogger(gomock.Any()).Return(nil).AnyTimes()
	mocked.Mailer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()
	mocked.TechnoRepository.EXPECT().FetchStack(context.Background()).Return(map[string]interface{}{}, nil).AnyTimes()

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
