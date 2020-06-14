package templater

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Templater interface {
	Render(name string, view interface{}, response http.ResponseWriter, status int)
}

type TemplateHtml struct {
	RootPath string
}

func (t *TemplateHtml) Render(name string, view interface{}, response http.ResponseWriter, status int) {
	response.WriteHeader(status)

	htmlTemplate := t.parseTemplates(name)
	err := htmlTemplate.ExecuteTemplate(response, name, view)

	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
}
func (t *TemplateHtml) parseTemplates(name string) *template.Template {
	templ := template.New(name)
	err := filepath.Walk(t.RootPath+"/tmpl/", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	return templ
}