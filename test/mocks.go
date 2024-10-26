package mocks

import (
	"fmt"

	contactMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/domain/repository"
	technoMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/domain/repository"
	mailerMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/infra/service/mailer"
	tracerMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/infra/tracer"

	"github.com/golang/mock/gomock"
)

type Mocks struct {
	ContactRepository *contactMock.MockContactRepository
	TechnoRepository  *technoMock.MockTechnoRepository
	HttpTracer        *tracerMock.MockTracer
	Mailer            *mailerMock.MockMailer
}

func New(reporter gomock.TestReporter) Mocks {
	controller := gomock.NewController(reporter)
	fmt.Println(controller)
	return Mocks{
		ContactRepository: contactMock.NewMockContactRepository(controller),
		HttpTracer:        tracerMock.NewMockTracer(controller),
		Mailer:            mailerMock.NewMockMailer(controller),
		TechnoRepository:  technoMock.NewMockTechnoRepository(controller),
	}
}
