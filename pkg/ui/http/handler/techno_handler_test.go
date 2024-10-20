package handler_test

import (
	"context"
	"github.com/Medzoner/medzoner-go/pkg/application/query"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/repository"
	"github.com/Medzoner/medzoner-go/pkg/ui/http/handler"
	tracerMock "github.com/Medzoner/medzoner-go/test/mocks/pkg/infra/tracer"
	"github.com/golang/mock/gomock"
	"go.opentelemetry.io/otel/trace/noop"
	"gotest.tools/assert"
	"net/http/httptest"
	"testing"
)

func TestTechnoHandler(t *testing.T) {
	t.Run("Unit: test TechnoHandler success", func(t *testing.T) {
		httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
		httpTracerMock.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).Times(5)
		httpTracerMock.EXPECT().WriteLog(gomock.Any(), gomock.Any()).Return().Times(1)
		technoHandler := handler.NewTechnoHandler(
			&TemplaterTest{},
			query.NewListTechnoQueryHandler(
				repository.NewTechnoJSONRepository(
					&LoggerTest{},
					&config.Config{
						RootPath: "./../../../../",
					},
				),
				httpTracerMock,
			),
			httpTracerMock,
		)

		responseWriter := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/techno", nil)
		technoHandler.IndexHandle(responseWriter, request)

		assert.Equal(t, responseWriter.Code, 200)
	})
	//
	//t.Run("Unit: test TechnoHandler with templater render failed", func(t *testing.T) {
	//	httpTracerMock := tracerMock.NewMockTracer(gomock.NewController(t))
	//	httpTracerMock.EXPECT().WriteLog(gomock.Any(), gomock.Any()).Return().Times(1)
	//	technoHandler := handler.NewTechnoHandler(
	//		&TemplaterTestFailed{},
	//		query.NewListTechnoQueryHandler(
	//			repository.NewTechnoJSONRepository(
	//				&LoggerTest{},
	//				&config.Config{
	//					RootPath: "./../../../../",
	//				},
	//			),
	//			httpTracerMock,
	//		),
	//		httpTracerMock,
	//	)
	//
	//	responseWriter := httptest.NewRecorder()
	//	request := httptest.NewRequest("Get", "/techno", nil)
	//	defer func() {
	//		if r := recover(); r == nil {
	//			t.Errorf("The code did not panic")
	//		}
	//	}()
	//	technoHandler.IndexHandle(responseWriter, request)
	//
	//	assert.Equal(t, responseWriter.Code, 500)
	//})
}
