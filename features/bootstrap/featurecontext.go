package bootstrap

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/Medzoner/medzoner-go/pkg"
	"github.com/Medzoner/medzoner-go/pkg/infra/database"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

//APIFeature APIFeature
type APIFeature struct {
	Response *http.Response
	Request  *http.Request
	BaseURL  *string
	App      *pkg.App
}

//BodyRequest BodyRequest
type BodyRequest struct {
	Body io.Reader
}

//Read Read
func (b BodyRequest) Read(p []byte) (n int, err error) {
	buffer := &bytes.Buffer{}
	return buffer.Read(p)
}

//New New
func New(url string, App *pkg.App) *APIFeature {
	feature := &APIFeature{Response: &http.Response{}, BaseURL: &url, App: App}
	feature.Request, _ = http.NewRequest("GET", fmt.Sprintf("%s%s", url, "/"), BodyRequest{}.Body)
	return feature
}

//FeatureContext FeatureContext
func (a *APIFeature) FeatureContext(s *godog.Suite) {
	s.BeforeSuite(func() {
		a.resetBdd()
	})
	s.BeforeScenario(func(interface{}) {
		a.resetResponse()
	})
	s.Step(`^I add "([^"]*)" header equal to "([^"]*)"$`, a.iAddHeaderEqualTo)
	s.Step(`^I send a GET request to "([^"]*)"$`, a.iSendAGETRequestTo)
	s.Step(`^I send a POST request to "([^"]*)" with body:$`, a.iSendAPOSTRequestToWithBody)
	s.Step(`^I send a PUT request to "([^"]*)" with body:$`, a.iSendAPUTRequestToWithBody)
	s.Step(`^I send a DELETE request to "([^"]*)"$`, a.iSendADELETERequestTo)
	s.Step(`^the response status code should be (\d+)$`, a.theResponseStatusCodeShouldBe)
	s.Step(`^the JSON node "([^"]*)" should be equal to "([^"]*)"$`, a.theJSONNodeShouldBeEqualTo)
	s.Step(`^the response should be in JSON$`, a.theResponseShouldBeInJSON)
	s.Step(`^the JSON should be valid according to the schema "([^"]*)"$`, a.theJSONShouldBeValidAccordingToTheSchema)
	s.Step(`^the JSON node "([^"]*)" should contain "([^"]*)"$`, a.theJSONNodeShouldContain)
	s.Step(`^the JSON node "([^"]*)" should be false$`, a.theJSONNodeShouldBeFalse)
	s.Step(`^the JSON node "([^"]*)" should not exist$`, a.theJSONNodeShouldNotExist)
	s.Step(`^the JSON node "([^"]*)" should contain (\d+)$`, a.theJSONNodeShouldContain)
	s.Step(`^the JSON node "([^"]*)" should have (\d+) elements$`, a.theJSONNodeShouldHaveElements)
	s.Step(`^PaginationScenario$`, a.paginationScenario)
	s.Step(`^print last JSON response$`, a.printLastJSONResponse)
	s.Step(`^the JSON node "([^"]*)" should be true$`, a.theJSONNodeShouldBeTrue)
	s.Step(`^the JSON node "([^"]*)" should exist$`, a.theJSONNodeShouldExist)
	s.Step(`^the JSON node "([^"]*)" should not be null$`, a.theJSONNodeShouldNotBeNull)
	s.Step(`^the JSON node "([^"]*)" should be null$`, a.theJSONNodeShouldBeNull)
	s.Step(`^the JSON node "([^"]*)" should contain \'The key "([^"]*)" is invalid as it will override the existing key "([^"]*)"\'$`, a.theJSONNodeShouldContainTheKeyIsInvalidAsItWillOverrideTheExistingKey)
	s.Step(`^the response should be empty$`, a.theResponseShouldBeEmpty)
}

func (a *APIFeature) resetResponse() {
	a.Request, _ = http.NewRequest("GET", "/", BodyRequest{}.Body)
	a.Response = &http.Response{}
}

func (a *APIFeature) iAddHeaderEqualTo(arg1 string, arg2 string) (err error) {
	a.Request.Header.Set(arg1, arg2)
	return
}

func (a *APIFeature) iSendARequestTo(method, endpoint string) (err error) {
	a.Request.Method = method
	a.Request.URL, err = url.Parse(fmt.Sprintf("%s%s", *a.BaseURL, endpoint))
	if err != nil {
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
	}()
	return
}

func (a *APIFeature) theResponseStatusCodeShouldBe(code int) (err error) {
	if code != a.Response.StatusCode && a.Response.Request.Response.StatusCode != code {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.Response.StatusCode)
	}
	return
}

func (a *APIFeature) theResponseShouldMatchJSON(body *gherkin.DocString) (err error) {
	var expected, actual []byte
	var data interface{}
	if err = json.Unmarshal([]byte(body.Content), &data); err != nil {
		return
	}
	if expected, err = json.Marshal(data); err != nil {
		return
	}
	actual, _ = ioutil.ReadAll(a.Response.Body)
	if !bytes.Equal(actual, expected) {
		err = fmt.Errorf("expected json, does not match actual: %s", string(actual))
	}
	return
}

func (a *APIFeature) theJSONNodeShouldBeEqualTo(arg1, arg2 string) (err error) {
	data := make(map[string]interface{})

	bodyBytes, err := ioutil.ReadAll(a.Response.Body)
	if err != nil {
		fmt.Println(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	if err = json.Unmarshal(bodyBytes, &data); err != nil {
		return
	}
	if arg2 != data[arg1] {
		err = fmt.Errorf("expected json, does not match actual: %s", arg1)
	}
	return
}

func (a *APIFeature) theResponseShouldBeInJSON() (err error) {
	res, _ := ioutil.ReadAll(a.Response.Body)
	var js json.RawMessage
	if json.Unmarshal(res, &js) != nil {
		return fmt.Errorf("expected response in json")
	}
	return nil
}

func (a *APIFeature) theJSONShouldBeValidAccordingToTheSchema(arg1 string) (err error) {
	_ = arg1
	return godog.ErrPending
}

func (a *APIFeature) theJSONNodeShouldContain(arg1, arg2 string) (err error) {
	_ = arg1
	_ = arg2
	return godog.ErrPending
}

func (a *APIFeature) theJSONNodeShouldBeFalse(arg1 string) (err error) {
	_ = arg1
	return godog.ErrPending
}

func (a *APIFeature) theJSONNodeShouldNotExist(arg1 string) (err error) {
	_ = arg1
	return godog.ErrPending
}

func (a *APIFeature) theJSONNodeShouldHaveElements(arg1 string, arg2 int) (err error) {
	_ = arg1
	_ = arg2
	return godog.ErrPending
}

func (a *APIFeature) paginationScenario() error {
	return godog.ErrPending
}

func (a *APIFeature) printLastJSONResponse() (err error) {
	bodyBytes, err := ioutil.ReadAll(a.Response.Body)
	if err != nil {
		fmt.Println(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	return
}

func (a *APIFeature) theJSONNodeShouldBeTrue(arg1 string) (err error) {
	_ = arg1
	return godog.ErrPending
}

func (a *APIFeature) theJSONNodeShouldExist(arg1 string) (err error) {
	_ = arg1
	return godog.ErrPending
}

func (a *APIFeature) theJSONNodeShouldNotBeNull(arg1 string) (err error) {
	_ = arg1
	return godog.ErrPending
}

func (a *APIFeature) theJSONNodeShouldBeNull(arg1 string) (err error) {
	_ = arg1
	return godog.ErrPending
}

func (a *APIFeature) theJSONNodeShouldContainTheKeyIsInvalidAsItWillOverrideTheExistingKey(arg1, arg2, arg3 string) (err error) {
	_ = arg1
	_ = arg2
	_ = arg3
	return godog.ErrPending
}

func (a *APIFeature) theResponseShouldBeEmpty() (err error) {
	bodyBytes, err := ioutil.ReadAll(a.Response.Body)
	if err != nil {
		fmt.Println(err)
	}
	bodyString := string(bodyBytes)
	if "" != bodyString {
		return fmt.Errorf("expected response body to be null, but actual is not empty")
	}
	return nil
}

func (a *APIFeature) iSendAGETRequestTo(arg1 string) (err error) {
	return a.iSendARequestTo("GET", arg1)
}

func (a *APIFeature) iSendADELETERequestTo(arg1 string) (err error) {
	return a.iSendARequestTo("DELETE", arg1)
}

func (a *APIFeature) iSendAPOSTRequestToWithBody(arg1 string, arg2 *gherkin.DocString) error {
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
	err := errors.New("err http send")
	a.Request.URL, err = url.Parse(fmt.Sprintf("%s%s", *a.BaseURL, arg1))
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.PostForm(a.Request.URL.String(), v)
	if err != nil {
		fmt.Println(err)
	}
	_ = resp.Body.Close()
	a.Response = resp

	// handle panic
	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()
	return nil
}

func (a *APIFeature) iSendAPUTRequestToWithBody(arg1 string, arg2 *gherkin.DocString) (err error) {
	return a.iSendAPOSTRequestToWithBody(arg1, arg2)
}

func (a *APIFeature) resetBdd() {
	a.App.Container.Get("database").(*database.DbSQLInstance).CreateDatabase()
	a.App.Container.Get("database").(*database.DbSQLInstance).DropDatabase()
	a.App.Container.Get("database").(*database.DbSQLInstance).CreateDatabase()
	a.App.Container.Get("db-manager").(*database.DbMigration).MigrateUp()
	return
}
