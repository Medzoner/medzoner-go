package main

import (
	"github.com/Medzoner/medzoner-go/features/bootstrap"
	"github.com/Medzoner/medzoner-go/pkg/infra/dependency"
	mocks "github.com/Medzoner/medzoner-go/test"
	tracerMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/infra/tracer"
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

	httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
	//httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(11)
	//httpTracerMock.EXPECT().Int64Counter(gomock.Any(), gomock.Any()).Return(metricNoop.Int64Counter{}, nil)
	//httpTracerMock.EXPECT().WriteLog(gomock.Any(), gomock.Any()).Return().Times(1)
	srv, err := dependency.InitServerTest(mockedRepository, *httpTracerMock)
	if err != nil {
		t.Error(err)
		return
	}
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
