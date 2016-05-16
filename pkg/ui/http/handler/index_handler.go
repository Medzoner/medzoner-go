package handler

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexHandler struct {
}

type IndexView struct {
	Locale    string
	PageTitle string
	TorHost   string
}

func (*IndexHandler) IndexHandle(w http.ResponseWriter, r *http.Request) {
	view := IndexView{Locale: "fr", PageTitle: "MedZoner.com"}

	view.TorHost = r.Header.Get("TOR-HOST")
	t := template.New("index template")
	pp, _ := os.Getwd()
	t = template.Must(t.ParseFiles(pp+"/tmpl/base.html", pp+"/tmpl/footer.html", pp+"/tmpl/header.html", pp+"/tmpl/Index/home.html"))
	w.WriteHeader(http.StatusOK)
	err := t.ExecuteTemplate(w, "layout", view)

	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
	_ = r
}
