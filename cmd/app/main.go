package main

import (
	"context"

	wiring "github.com/Medzoner/medzoner-go/pkg/infra/dependency"
)

func main() {
	ctx := context.Background()

	server, err := wiring.InitServer()
	if err != nil {
		panic(err)
	}

	server.Start(ctx)
}
