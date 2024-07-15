package main

import (
	"context"
	"github.com/jotadrilo/go-factory/cmd"
	"github.com/jotadrilo/go-factory/pkg/log"
)

func main() {
	var ctx = context.Background()

	if err := mainE(ctx); err != nil {
		log.Logger.Fatalf("Process failed: %s", err.Error())
	}
}

func mainE(ctx context.Context) error {
	return cmd.NewRootCmd().ExecuteContext(ctx)
}
