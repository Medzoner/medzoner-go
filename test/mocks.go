package mocks

import (
	"fmt"
	contactMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/domain/repository"
	tracerMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/infra/tracer"

	"github.com/golang/mock/gomock"
)

type Mocks struct {
	ContactRepository *contactMock.MockContactRepository
	HttpTracer        *tracerMock.MockTracer
}

func New(reporter gomock.TestReporter) Mocks {
	controller := gomock.NewController(reporter)
	fmt.Println(controller)
	return Mocks{
		ContactRepository: contactMock.NewMockContactRepository(controller),
		HttpTracer:        tracerMock.NewMockTracer(controller),
	}
}
