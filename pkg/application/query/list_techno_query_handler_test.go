package query_test

import (
	"context"
	"errors"
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	mocks "github.com/Medzoner/medzoner-go/test"
	tracerMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/infra/telemetry"
	"github.com/golang/mock/gomock"
	"go.opentelemetry.io/otel/trace/noop"
	"testing"
)

func TestListTechnoQueryHandler(t *testing.T) {
	t.Run("Unit: test ListTechnoQueryHandler \"stack\" success", func(t *testing.T) {
		listTechnoQuery := query.ListTechnoQuery{
			Type: "stack",
		}
		httpTelemetryMock := tracerMock.NewMockTelemeter(gomock.NewController(t))
		httpTelemetryMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)

		handler := query.NewListTechnoQueryHandler(&TechnoRepositoryTest{}, httpTelemetryMock)

		_, err := handler.Handle(context.Background(), listTechnoQuery)
		if err != nil {
			t.Errorf("ErrorSpan: %v", err)
		}
	})

	t.Run("Unit: test ListTechnoQueryHandler \"experience\" success", func(t *testing.T) {
		listTechnoQuery := query.ListTechnoQuery{
			Type: "experience",
		}
		httpTelemetryMock := tracerMock.NewMockTelemeter(gomock.NewController(t))
		httpTelemetryMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		handler := query.NewListTechnoQueryHandler(&TechnoRepositoryTest{}, httpTelemetryMock)

		_, err := handler.Handle(context.Background(), listTechnoQuery)
		if err != nil {
			t.Errorf("ErrorSpan: %v", err)
		}
	})

	t.Run("Unit: test ListTechnoQueryHandler \"formation\" success", func(t *testing.T) {
		listTechnoQuery := query.ListTechnoQuery{
			Type: "formation",
		}
		httpTelemetryMock := tracerMock.NewMockTelemeter(gomock.NewController(t))
		httpTelemetryMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		handler := query.NewListTechnoQueryHandler(&TechnoRepositoryTest{}, httpTelemetryMock)

		_, err := handler.Handle(context.Background(), listTechnoQuery)
		if err != nil {
			t.Errorf("ErrorSpan: %v", err)
		}
	})

	t.Run("Unit: test ListTechnoQueryHandler \"lang\" success", func(t *testing.T) {
		listTechnoQuery := query.ListTechnoQuery{
			Type: "lang",
		}
		httpTelemetryMock := tracerMock.NewMockTelemeter(gomock.NewController(t))
		httpTelemetryMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		handler := query.NewListTechnoQueryHandler(&TechnoRepositoryTest{}, httpTelemetryMock)

		_, err := handler.Handle(context.Background(), listTechnoQuery)
		if err != nil {
			t.Errorf("ErrorSpan: %v", err)
		}
	})

	t.Run("Unit: test ListTechnoQueryHandler \"other\" success", func(t *testing.T) {
		listTechnoQuery := query.ListTechnoQuery{
			Type: "other",
		}
		httpTelemetryMock := tracerMock.NewMockTelemeter(gomock.NewController(t))
		httpTelemetryMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		handler := query.NewListTechnoQueryHandler(&TechnoRepositoryTest{}, httpTelemetryMock)

		_, err := handler.Handle(context.Background(), listTechnoQuery)
		if err != nil {
			t.Errorf("ErrorSpan: %v", err)
		}
	})

	t.Run("Unit: test ListTechnoQueryHandler non existent type success", func(t *testing.T) {
		listTechnoQuery := query.ListTechnoQuery{
			Type: "fake",
		}
		httpTelemetryMock := tracerMock.NewMockTelemeter(gomock.NewController(t))
		httpTelemetryMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		handler := query.NewListTechnoQueryHandler(&TechnoRepositoryTest{}, httpTelemetryMock)

		_, err := handler.Handle(context.Background(), listTechnoQuery)
		if err != nil {
			t.Errorf("ErrorSpan: %v", err)
		}
	})

	t.Run("Unit: test ListTechnoQueryHandler \"stack\" failed", func(t *testing.T) {
		mocked := mocks.New(t)
		mocked.TechnoRepository.EXPECT().FetchStack(gomock.Any()).Return(nil, errors.New("error")).AnyTimes()
		listTechnoQuery := query.ListTechnoQuery{
			Type: "stack",
		}
		httpTelemetryMock := tracerMock.NewMockTelemeter(gomock.NewController(t))
		httpTelemetryMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		httpTelemetryMock.EXPECT().ErrorSpan(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

		handler := query.NewListTechnoQueryHandler(mocked.TechnoRepository, httpTelemetryMock)

		_, err := handler.Handle(context.Background(), listTechnoQuery)
		if err != nil {
			t.Errorf("ErrorSpan: %v", err)
		}
	})
}

type TechnoRepositoryTest struct{}

func (m *TechnoRepositoryTest) FetchStack(ctx context.Context) (map[string]interface{}, error) {
	_ = ctx
	return map[string]interface{}{}, nil
}
