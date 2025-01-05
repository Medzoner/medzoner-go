package mocks

import (
	"fmt"

	domainRepository "github.com/Medzoner/medzoner-go/test/mocks/pkg/domain/repository"
	mailerMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/infra/service/mailer"
	tracerMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/infra/telemetry"

	"github.com/golang/mock/gomock"
)

type Mocks struct {
	ContactRepository *domainRepository.MockContactRepository
	TechnoRepository  *domainRepository.MockTechnoRepository
	HttpTelemetry     *tracerMock.MockTelemeter
	Mailer            *mailerMock.MockMailer
}

func New(reporter gomock.TestReporter) Mocks {
	controller := gomock.NewController(reporter)
	fmt.Println(controller)
	return Mocks{
		ContactRepository: domainRepository.NewMockContactRepository(controller),
		HttpTelemetry:     tracerMock.NewMockTelemeter(controller),
		Mailer:            mailerMock.NewMockMailer(controller),
		TechnoRepository:  domainRepository.NewMockTechnoRepository(controller),
	}
}
