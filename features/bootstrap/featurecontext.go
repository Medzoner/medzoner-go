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

	"github.com/Medzoner/medzoner-go/pkg/infra/server"
	mocks "github.com/Medzoner/medzoner-go/test"

	"github.com/cucumber/godog"
)

// APIFeature APIFeature
type APIFeature struct {
	Response *http.Response
	Request  *http.Request
	Server   server.Server
	Mocks    mocks.Mocks
}

// BodyRequest BodyRequest
type BodyRequest struct {
	Body io.Reader
}

// Read implement io.Reader
func (b BodyRequest) Read(p []byte) (n int, err error) {
	buffer := &bytes.Buffer{}
	return buffer.Read(p)
}

// New initialize a new APIFeature
func New(srv server.Server, mocked mocks.Mocks) *APIFeature {
	feature := &APIFeature{
		Response: &http.Response{},
		Server:   srv,
		Mocks:    mocked,
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", recorder.Body)
	feature.Request = request

	srv.Router.ServeHTTP(recorder, request)
	return feature
}

// InitializeTestSuite InitializeTestSuite
func (a *APIFeature) InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		fmt.Println("BeforeSuite", ctx)
		//a.resetBdd()
	})
	ctx.AfterSuite(func() {
		fmt.Println("AfterSuite", ctx)
		//mg := wiring.InitDbMigration()
		//mg.MigrateDown()
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
	a.Request, _ = http.NewRequest("GET", "/", BodyRequest{}.Body)
	a.Response = &http.Response{}
}

func (a *APIFeature) iAddHeaderEqualTo(arg1, arg2 string) (err error) {
	a.Request.Header.Set(arg1, arg2)
	return
}

func (a *APIFeature) iSendARequestTo(method, endpoint string) (err error) {
	a.Request.Method = method
	a.Request.URL, err = url.Parse(endpoint)
	/*	if err != nil {
			return
		}

		client := &http.Client{}
		resp, err := client.Do(a.Request)
		if err != nil {
			fmt.Println(err)
		}
		// _ = resp.Body.Close()
		a.Response = resp

		// handle panic
		defer func() {
			switch t := recover().(type) {
			case string:
				err = fmt.Errorf(t)
			case error:
				err = t
			}
		}()*/

	recorder := httptest.NewRecorder()

	request := a.Request
	a.Server.Router.ServeHTTP(recorder, request)
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
	a.Request.Method = "POST"
	urlParse, err := url.Parse(arg1)
	if err != nil {
		return err
	}
	a.Request.URL = urlParse

	recorder := httptest.NewRecorder()
	request := a.Request

	a.Server.Router.ServeHTTP(recorder, request)
	a.Response = recorder.Result()
	return nil
}
