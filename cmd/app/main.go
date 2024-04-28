package main

import (
	wiring "github.com/Medzoner/medzoner-go/pkg/infra/dependency"
)

func main() {
	//rootPath, _ := os.Getwd()
	app := wiring.InitApp()
	//builder, _ := di.NewBuilder()
	//appli.LoadContainer(builder)
	app.Handle("web")
}
