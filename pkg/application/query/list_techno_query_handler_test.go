package query_test

import (
	"context"
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	tracerMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/infra/tracer"
	"github.com/golang/mock/gomock"
	"go.opentelemetry.io/otel/trace/noop"
	"testing"
)

func TestListTechnoQueryHandler(t *testing.T) {

	t.Run("Unit: test ListTechnoQueryHandler \"stack\" success", func(t *testing.T) {
		listTechnoQuery := query.ListTechnoQuery{
			Type: "stack",
		}
		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)

		handler := query.NewListTechnoQueryHandler(&TechnoRepositoryTest{}, httpTracerMock)

		_, err := handler.Handle(context.Background(), listTechnoQuery)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("Unit: test ListTechnoQueryHandler \"experience\" success", func(t *testing.T) {
		listTechnoQuery := query.ListTechnoQuery{
			Type: "experience",
		}
		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		handler := query.NewListTechnoQueryHandler(&TechnoRepositoryTest{}, httpTracerMock)

		_, err := handler.Handle(context.Background(), listTechnoQuery)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("Unit: test ListTechnoQueryHandler \"formation\" success", func(t *testing.T) {
		listTechnoQuery := query.ListTechnoQuery{
			Type: "formation",
		}
		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		handler := query.NewListTechnoQueryHandler(&TechnoRepositoryTest{}, httpTracerMock)

		_, err := handler.Handle(context.Background(), listTechnoQuery)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("Unit: test ListTechnoQueryHandler \"lang\" success", func(t *testing.T) {
		listTechnoQuery := query.ListTechnoQuery{
			Type: "lang",
		}
		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		handler := query.NewListTechnoQueryHandler(&TechnoRepositoryTest{}, httpTracerMock)

		_, err := handler.Handle(context.Background(), listTechnoQuery)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("Unit: test ListTechnoQueryHandler \"other\" success", func(t *testing.T) {
		listTechnoQuery := query.ListTechnoQuery{
			Type: "other",
		}
		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		handler := query.NewListTechnoQueryHandler(&TechnoRepositoryTest{}, httpTracerMock)

		_, err := handler.Handle(context.Background(), listTechnoQuery)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})

	t.Run("Unit: test ListTechnoQueryHandler non existent type success", func(t *testing.T) {
		listTechnoQuery := query.ListTechnoQuery{
			Type: "fake",
		}
		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(1)
		handler := query.NewListTechnoQueryHandler(&TechnoRepositoryTest{}, httpTracerMock)

		_, err := handler.Handle(context.Background(), listTechnoQuery)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})
}

type TechnoRepositoryTest struct{}

func (m *TechnoRepositoryTest) FetchStack(ctx context.Context) (map[string]interface{}, error) {
	_ = ctx
	return map[string]interface{}{}, nil
}
