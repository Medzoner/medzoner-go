package templater

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Medzoner/medzoner-go/pkg/infra/config"
)

// Templater is an interface for rendering templates
type Templater interface {
	Render(name string, view interface{}, response http.ResponseWriter, status int) (interface{}, error)
}

// TemplateHTML is a struct that implements Templater interface
type TemplateHTML struct {
	RootPath string
}

// NewTemplateHTML is a function that returns a new TemplateHTML
func NewTemplateHTML(config config.Config) *TemplateHTML {
	return &TemplateHTML{
		RootPath: string(config.RootPath),
	}
}

// Render renders template
func (t *TemplateHTML) Render(name string, view interface{}, response http.ResponseWriter, status int) (interface{}, error) {
	_ = status

	htmlTemplate, err := t.parseTemplates(name)
	if err != nil {
		return nil, fmt.Errorf("error parsing templates: %v", err)
	}

	return nil, htmlTemplate.ExecuteTemplate(response, name, view)
}

// parseTemplates parses templates
func (t *TemplateHTML) parseTemplates(name string) (*template.Template, error) {
	tpl := template.New(name)
	err := filepath.Walk(t.RootPath+"tmpl/", func(path string, info os.FileInfo, err error) error {
		_ = info
		if strings.Contains(path, ".html") {
			_, err = tpl.ParseFiles(path)
			if err != nil {
				return fmt.Errorf("error parsing files tpl: %v", err)
			}
		}
		return err
	})

	return tpl, err
}
