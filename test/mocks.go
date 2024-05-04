package mocks

import (
	"fmt"
	contactMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/domain/repository"

	"github.com/golang/mock/gomock"
)

type Mocks struct {
	ContactRepository *contactMock.MockContactRepository
}

func New(reporter gomock.TestReporter) Mocks {
	controller := gomock.NewController(reporter)
	fmt.Println(controller)
	return Mocks{
		ContactRepository: contactMock.NewMockContactRepository(controller),
	}
}
