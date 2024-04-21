package templater_test

import (
	"github.com/Medzoner/medzoner-go/pkg/ui/http/templater"
	"gotest.tools/assert"
	"log"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRender(t *testing.T) {
	t.Run("Unit: test Render success", func(t *testing.T) {
		var tpl = templater.TemplateHTML{
			RootPath: "../../../..",
		}
		_, err := tpl.Render(
			"index",
			nil,
			httptest.NewRecorder(),
			200,
		)
		if err != nil {
			assert.Equal(t, true, false)
		}
		if err == nil {
			assert.Equal(t, true, true)
		}
	})
	t.Run("Unit: test Render failed when parse tpl", func(t *testing.T) {
		var tpl = templater.TemplateHTML{
			RootPath: "../../../../.var/test",
		}
		err := os.Chmod(tpl.RootPath+"/tmpl/failed.html", 0000)
		if err != nil {
			log.Fatal(err)
		}

		_, err = tpl.Render(
			"failed.html",
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
		err = os.Chmod(tpl.RootPath+"/tmpl/failed.html", 0700)
		if err != nil {
			log.Fatal(err)
		}
	})
	t.Run("Unit: test Render failed with bad tpl name", func(t *testing.T) {
		var tpl = templater.TemplateHTML{
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
	t.Run("Unit: test Render failed with bad rootPath", func(t *testing.T) {
		var tpl = templater.TemplateHTML{
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
