package main

import (
	"context"
	"github.com/jotadrilo/go-factory/cmd"
	"log"
)

func main() {
	var ctx = context.Background()

	if err := mainE(ctx); err != nil {
		log.Fatalf("Process failed: %v", err)
	}
}

func mainE(ctx context.Context) error {
	return cmd.NewRootCmd().ExecuteContext(ctx)
}
