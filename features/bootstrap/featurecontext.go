package bootstrap

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/Medzoner/gomedz/pkg/http/server"
	mocks "github.com/Medzoner/medzoner-go/test"
	"github.com/cucumber/godog"
)

// APIFeature APIFeature
type APIFeature struct {
	Mocks    mocks.Mocks
	Response *http.Response
	Request  *http.Request
	Server   server.Server
}

// BodyRequest BodyRequest
type BodyRequest struct {
	Body io.Reader
}

// Read implement io.Reader
func (b BodyRequest) Read(p []byte) (n int, err error) {
	_ = b
	buffer := &bytes.Buffer{}
	i, err := buffer.Read(p)
	if err != nil {
		return 0, fmt.Errorf("error reading body: %w", err)
	}
	return i, nil
}

// New initialize a new APIFeature
func New(srv server.Server, mocked mocks.Mocks) *APIFeature {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", recorder.Body)

	err := srv.Serve(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	return &APIFeature{
		Response: &http.Response{},
		Request:  request,
		Server:   srv,
		Mocks:    mocked,
	}
}

// InitializeTestSuite InitializeTestSuite
func (a *APIFeature) InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		// a.resetBdd()
	})
	ctx.AfterSuite(func() {
		if err := a.Server.Shutdown(context.Background()); err != nil {
			fmt.Println(err)
		}
	})
}

// InitializeScenario InitializeScenario
func (a *APIFeature) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, s *godog.Scenario) (context.Context, error) {
		_ = s
		a.resetResponse()
		return ctx, nil
	})
	ctx.Step(`^I add "([^"]*)" header equal to "([^"]*)"$`, a.iAddHeaderEqualTo)
	ctx.Step(`^I send a GET request to "([^"]*)"$`, a.iSendAGETRequestTo)
	ctx.Step(`^I send a POST request to "([^"]*)" with body:$`, a.iSendAPOSTRequestToWithBody)
	ctx.Step(`^the response status code should be (\d+)$`, a.theResponseStatusCodeShouldBe)
}

func (a *APIFeature) resetResponse() {
	a.Request, _ = http.NewRequest(http.MethodGet, "/", BodyRequest{}.Body)
	a.Response = &http.Response{}
}

func (a *APIFeature) iAddHeaderEqualTo(arg1, arg2 string) (err error) {
	a.Request.Header.Set(arg1, arg2)
	return
}

func (a *APIFeature) iSendARequestTo(method, endpoint string) (err error) {
	a.Request.Method = method
	a.Request.URL, err = url.Parse(endpoint)

	recorder := httptest.NewRecorder()
	a.Server.Serve(context.Background())
	a.Response = recorder.Result()

	return
}

func (a *APIFeature) theResponseStatusCodeShouldBe(code int) (err error) {
	if code < http.StatusBadRequest || code >= http.StatusInternalServerError {
		if code != a.Response.StatusCode || (a.Response.Request != nil && a.Response.Request.Response.StatusCode != code) {
			return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.Response.StatusCode)
		}
	}
	if a.Response.Request != nil && a.Response.Request.Response != nil && a.Response.Request.Response.StatusCode != code {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.Response.StatusCode)
	}
	return
}

func (a *APIFeature) iSendAGETRequestTo(arg1 string) (err error) {
	return a.iSendARequestTo("GET", arg1)
}

func (a *APIFeature) iSendAPOSTRequestToWithBody(arg1 string, arg2 *godog.DocString) error {
	v := url.Values{}
	if arg2 != nil && arg2.Content != "" {
		var data map[string]string
		if errUnmarshal := json.Unmarshal([]byte(arg2.Content), &data); errUnmarshal != nil {
			fmt.Println(errUnmarshal)
		}
		for key, value := range data {
			v.Set(key, value)
		}
		a.Request.PostForm = v
	}

	urlParse, err := url.Parse(arg1)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	a.Request.URL = urlParse
	a.Request.Method = http.MethodPost

	recorder := httptest.NewRecorder()

	a.Server.Serve(context.Background())
	a.Response = recorder.Result()

	return nil
}
