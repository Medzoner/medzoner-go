package main

import (
	"context"

	"github.com/Medzoner/gomedz/pkg/logger"
	wire "github.com/Medzoner/medzoner-go/pkg/dependency"
)

func main() {
	ctx := context.Background()
	srv, err := wire.InitServer(ctx)
	if err != nil {
		logger.Fatal(ctx, "Failed to initialize server", err)
		return
	}

	if err := srv.Serve(ctx); err != nil {
		logger.Fatal(ctx, "Server encountered an error", err)
	}
}
