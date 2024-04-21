package templater

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Templater Templater
type Templater interface {
	Render(name string, view interface{}, response http.ResponseWriter, status int) (interface{}, error)
}

// TemplateHTML TemplateHTML
type TemplateHTML struct {
	RootPath string
}

// NewTemplateHTML NewTemplateHTML
func NewTemplateHTML(rootPath string) *TemplateHTML {
	return &TemplateHTML{
		RootPath: rootPath,
	}
}

// Render Render
func (t *TemplateHTML) Render(name string, view interface{}, response http.ResponseWriter, status int) (interface{}, error) {
	response.WriteHeader(status)

	htmlTemplate, err := t.parseTemplates(name)
	if err != nil {
		return nil, err
	}
	err = htmlTemplate.ExecuteTemplate(response, name, view)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (t *TemplateHTML) parseTemplates(name string) (*template.Template, error) {
	templ := template.New(name)
	err := filepath.Walk(t.RootPath+"/tmpl/", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return templ, nil
}
