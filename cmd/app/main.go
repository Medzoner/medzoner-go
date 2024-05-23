package main

import (
	wiring "github.com/Medzoner/medzoner-go/pkg/infra/dependency"
)

func main() {
	server, err := wiring.InitServer()
	if err != nil {
		panic(err)
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
