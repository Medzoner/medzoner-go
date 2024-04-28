package main

import (
	wiring "github.com/Medzoner/medzoner-go/pkg/infra/dependency"
)

func main() {
	app := wiring.InitApp()
	app.Handle("migrate")
}
