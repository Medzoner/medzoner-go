package main

import (
	"github.com/Medzoner/medzoner-go/pkg/app"
	"github.com/Medzoner/medzoner-go/pkg/infra/path"
	"github.com/sarulabs/di"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	app := app.App{
		RootPath: path.RootPath(rootPath),
	}
	builder, _ := di.NewBuilder()
	app.LoadContainer(builder)
	app.Handle("migrate")
}
