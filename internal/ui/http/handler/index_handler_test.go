package handler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Medzoner/medzoner-go/pkg/infra/dependency"
	"github.com/Medzoner/medzoner-go/pkg/infra/server"
	mocks "github.com/Medzoner/medzoner-go/test"

	"github.com/golang/mock/gomock"
	"go.opentelemetry.io/otel/trace/noop"
	"gotest.tools/assert"
)

func TestIntegration_IndexHandler_Success(t *testing.T) {
	mocked := mocks.New(t)
	mocked.HttpTelemetry.EXPECT().StartRoot(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
	mocked.HttpTelemetry.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
	mocked.HttpTelemetry.EXPECT().ShutdownTracer(gomock.Any()).Return(nil).AnyTimes()
	mocked.HttpTelemetry.EXPECT().ShutdownMeter(gomock.Any()).Return(nil).AnyTimes()
	mocked.HttpTelemetry.EXPECT().ShutdownLogger(gomock.Any()).Return(nil).AnyTimes()
	mocked.HttpTelemetry.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
	mocked.HttpTelemetry.EXPECT().Log(gomock.Any(), gomock.Any()).AnyTimes()
	mocked.Mailer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()
	t.Setenv("APP_ENV", "test")
	t.Setenv("DEBUG", "true")
	t.Setenv("ROOT_PATH", "./../../../../")
	srv, err := dependency.InitServerTest(&mocked)
	if err != nil {
		t.Error(err)
	}
	defer func(srv *server.Server) {
		if err := srv.Shutdown(context.Background()); err != nil {
			t.Error(err)
		}
	}(srv)

	testCase := []struct {
		name         string
		method       string
		url          string
		body         url.Values
		expectedCode int
		mocks        func()
	}{
		{
			name:   "Unit: GET test IndexHandler success",
			method: http.MethodGet,
			url:    "/",
			body:   url.Values{},
			mocks: func() {
				mocked.TechnoRepository.EXPECT().FetchStack(context.Background()).Return(map[string]interface{}{}, nil).Times(1)
			},
			expectedCode: http.StatusOK,
		},
		{
			name:   "Unit: POST test IndexHandler with form submit success",
			method: http.MethodPost,
			url:    "/",
			body: url.Values{
				"name":               []string{"a name"},
				"email":              []string{"fake@fake.fake"},
				"message":            []string{"a message"},
				"g-captcha-response": []string{"captcha"},
				"submit":             []string{""},
			},
			mocks: func() {
				mocked.TechnoRepository.EXPECT().FetchStack(context.Background()).Return(map[string]interface{}{}, nil).Times(1)
				mocked.ContactRepository.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedCode: http.StatusSeeOther,
		},
		{
			name:   "Unit: test IndexHandler with form submit failed on struct",
			method: http.MethodPost,
			url:    "/",
			body: url.Values{
				"name":               []string{"a name"},
				"email":              []string{""},
				"message":            []string{"a message"},
				"g-captcha-response": []string{"captcha"},
				"submit":             []string{""},
			},
			mocks: func() {
				mocked.TechnoRepository.EXPECT().FetchStack(context.Background()).Return(map[string]interface{}{}, nil).Times(1)
			},
			expectedCode: http.StatusOK,
		},
		{
			name:   "Unit: test IndexHandler with TechnoRepository save failed",
			method: http.MethodPost,
			url:    "/",
			body: url.Values{
				"name":               []string{"a name"},
				"email":              []string{"fake@fakem.lan"},
				"message":            []string{"a message"},
				"g-captcha-response": []string{"captcha"},
				"submit":             []string{""},
			},
			mocks: func() {
				mocked.TechnoRepository.EXPECT().FetchStack(context.Background()).Return(map[string]interface{}{}, nil).Times(1)
				mocked.ContactRepository.EXPECT().Save(gomock.Any(), gomock.Any()).Return(errors.New("error")).Times(1)
				mocked.HttpTelemetry.EXPECT().ErrorSpan(gomock.Any(), gomock.Any()).Return(errors.New("error")).Times(1)
			},
			expectedCode: http.StatusInternalServerError,
		},
		{
			name:   "Unit: test IndexHandler with list techno failed",
			method: http.MethodGet,
			url:    "/",
			body:   url.Values{},
			mocks: func() {
				mocked.TechnoRepository.EXPECT().FetchStack(context.Background()).Return(nil, errors.New("error")).Times(1)
				mocked.HttpTelemetry.EXPECT().ErrorSpan(gomock.Any(), gomock.Any()).Return(errors.New("error")).Times(1)
			},
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			tc.mocks()
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(tc.method, tc.url, nil)
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			request.Header.Set("X-Correlation-ID", "test")
			request.Form = tc.body
			srv.Router.ServeHTTP(recorder, request)

			assert.Equal(t, recorder.Code, tc.expectedCode)
		})
	}
}

func TestIntegration_IndexHandler_Failed_Tpl(t *testing.T) {
	mocked := mocks.New(t)
	mocked.HttpTelemetry.EXPECT().StartRoot(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
	mocked.HttpTelemetry.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
	mocked.HttpTelemetry.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
	mocked.HttpTelemetry.EXPECT().Log(gomock.Any(), gomock.Any()).AnyTimes()
	mocked.TechnoRepository.EXPECT().FetchStack(context.Background()).Return(map[string]interface{}{}, nil).AnyTimes()
	t.Setenv("APP_ENV", "test")
	t.Setenv("DEBUG", "true")
	t.Setenv("ROOT_PATH", "./")
	srv, err := dependency.InitServerTest(&mocked)
	if err != nil {
		t.Error(err)
	}

	t.Run("Unit: GET test IndexHandler failed template path", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Form = url.Values{}
		srv.Router.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusInternalServerError)
	})
}

func TestIntegration_IndexHandler_Failed_Captcha(t *testing.T) {
	mocked := mocks.New(t)
	mocked.HttpTelemetry.EXPECT().StartRoot(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
	mocked.HttpTelemetry.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
	mocked.TechnoRepository.EXPECT().FetchStack(context.Background()).Return(map[string]interface{}{}, nil).AnyTimes()
	t.Setenv("APP_ENV", "test")
	t.Setenv("DEBUG", "false") // to avoid error on recaptcha
	t.Setenv("ROOT_PATH", "./../../../../")
	srv, err := dependency.InitServerTest(&mocked)
	if err != nil {
		t.Error(err)
	}

	t.Run("Unit: POST test IndexHandler failed captcha", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/", nil)
		request.Form = url.Values{
			"g-captcha-response": []string{"captcha"},
		}
		srv.Router.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusSeeOther)
	})
}
