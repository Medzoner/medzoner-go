package main

import (
	"github.com/Medzoner/medzoner-go/pkg"
)

func main() {
	App := pkg.App{}
	App.Handle("web")
}
