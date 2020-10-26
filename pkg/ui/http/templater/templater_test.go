package templater_test

import (
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"gotest.tools/assert"
	"net/http/httptest"
	"testing"
)

func TestRender(t *testing.T) {
	t.Run("Unit: test Render failed", func(t *testing.T) {
		var tpl = templater.TemplateHtml{
			RootPath: "../../../..",
		}
		_, err := tpl.Render(
			"fail",
			nil,
			httptest.NewRecorder(),
			200,
		)
		if err != nil {
			assert.Equal(t, true, true)
		}
		if err == nil {
			assert.Equal(t, true, false)
		}
	})
	t.Run("Unit: test Render failed", func(t *testing.T) {
		var tpl = templater.TemplateHtml{
			RootPath: "../../..",
		}
		_, err := tpl.Render(
			"index",
			nil,
			httptest.NewRecorder(),
			200,
		)
		if err != nil {
			assert.Equal(t, true, true)
		}
		if err == nil {
			assert.Equal(t, true, false)
		}
	})
}
