package main

import (
	"github.com/Medzoner/medzoner-go/pkg"
	"github.com/sarulabs/di"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	app := pkg.App{
		RootPath: rootPath,
	}
	builder, _ := di.NewBuilder()
	app.LoadContainer(builder)
	app.Handle("web")
}
