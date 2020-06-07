package main

import (
	"github.com/Medzoner/medzoner-go/pkg"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	App := pkg.App{
		RootPath: rootPath,
	}
	App.Handle("web")
}
