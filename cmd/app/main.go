package main

import (
	wiring "github.com/Medzoner/medzoner-go/pkg/infra/dependency"
)

func main() {
	server := wiring.InitServer()
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
