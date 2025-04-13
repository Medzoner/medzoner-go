package mocks

import (
	"fmt"

	"github.com/Medzoner/medzoner-go/test/mocks"
	"github.com/golang/mock/gomock"
)

type Mocks struct {
	ContactRepository *mocks.MockContactRepository
	TechnoRepository  *mocks.MockTechnoRepository
	HttpTelemetry     *mocks.MockTelemeter
	Mailer            *mocks.MockMailer
}

func New(reporter gomock.TestReporter) Mocks {
	controller := gomock.NewController(reporter)
	fmt.Println(controller)
	return Mocks{
		ContactRepository: mocks.NewMockContactRepository(controller),
		HttpTelemetry:     mocks.NewMockTelemeter(controller),
		Mailer:            mocks.NewMockMailer(controller),
		TechnoRepository:  mocks.NewMockTechnoRepository(controller),
	}
}
