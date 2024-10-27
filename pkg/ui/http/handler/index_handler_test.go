package handler_test

import (
	"context"
	"errors"
	"github.com/Medzoner/medzoner-go/pkg/infra/server"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/Medzoner/medzoner-go/pkg/infra/dependency"
	mocks "github.com/Medzoner/medzoner-go/test"
	"github.com/golang/mock/gomock"
	"go.opentelemetry.io/otel/trace/noop"
	"gotest.tools/assert"
)

func TestIntegration_IndexHandler_Success(t *testing.T) {
	mocked := mocks.New(t)
	mocked.HttpTracer.EXPECT().StartRoot(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
	mocked.HttpTracer.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
	mocked.HttpTracer.EXPECT().ShutdownTracer(gomock.Any()).Return(nil).AnyTimes()
	mocked.HttpTracer.EXPECT().ShutdownMeter(gomock.Any()).Return(nil).AnyTimes()
	mocked.HttpTracer.EXPECT().ShutdownLogger(gomock.Any()).Return(nil).AnyTimes()
	mocked.Mailer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()
	_ = os.Setenv("APP_ENV", "test")
	_ = os.Setenv("DEBUG", "true")
	_ = os.Setenv("ROOT_PATH", "./../../../../")
	srv, err := dependency.InitServerTest(&mocked)
	if err != nil {
		t.Error(err)
	}
	defer func(srv *server.Server, ctx context.Context) {
		if err := srv.Shutdown(ctx); err != nil {
			t.Error(err)
		}
	}(srv, context.Background())

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
				"name":               {"a name"},
				"email":              {"fake@fake.fake"},
				"message":            {"a message"},
				"g-captcha-response": {"captcha"},
				"submit":             {""},
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
				"name":               {"a name"},
				"email":              {""},
				"message":            {"a message"},
				"g-captcha-response": {"captcha"},
				"submit":             {""},
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
				"name":               {"a name"},
				"email":              {"fake@fakem.lan"},
				"message":            {"a message"},
				"g-captcha-response": {"captcha"},
				"submit":             {""},
			},
			mocks: func() {
				mocked.TechnoRepository.EXPECT().FetchStack(context.Background()).Return(map[string]interface{}{}, nil).Times(1)
				mocked.ContactRepository.EXPECT().Save(gomock.Any(), gomock.Any()).Return(errors.New("error")).Times(1)
				mocked.HttpTracer.EXPECT().Error(gomock.Any(), gomock.Any()).Return(errors.New("error")).Times(1)
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
				mocked.HttpTracer.EXPECT().Error(gomock.Any(), gomock.Any()).Return(errors.New("error")).Times(1)
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
	mocked.HttpTracer.EXPECT().StartRoot(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
	mocked.HttpTracer.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
	mocked.TechnoRepository.EXPECT().FetchStack(context.Background()).Return(map[string]interface{}{}, nil).AnyTimes()
	_ = os.Setenv("APP_ENV", "test")
	_ = os.Setenv("DEBUG", "true")
	_ = os.Setenv("ROOT_PATH", "./")
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
	mocked.HttpTracer.EXPECT().StartRoot(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
	mocked.HttpTracer.EXPECT().Start(gomock.Any(), gomock.Any(), gomock.Any()).Return(context.Background(), noop.Span{}).AnyTimes()
	mocked.TechnoRepository.EXPECT().FetchStack(context.Background()).Return(map[string]interface{}{}, nil).AnyTimes()
	_ = os.Setenv("APP_ENV", "test")
	_ = os.Setenv("DEBUG", "false") // to avoid error on recaptcha
	_ = os.Setenv("ROOT_PATH", "./../../../../")
	srv, err := dependency.InitServerTest(&mocked)
	if err != nil {
		t.Error(err)
	}

	t.Run("Unit: POST test IndexHandler failed captcha", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/", nil)
		request.Form = url.Values{
			"g-captcha-response": {"captcha"},
		}
		srv.Router.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusSeeOther)
	})
}
