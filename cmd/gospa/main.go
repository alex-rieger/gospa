package main

import (
	"fmt"
	"os"

	"github.com/alex-rieger/gospa/internal"
	"go.uber.org/zap"

	// This controls the maxprocs environment variable in container runtimes.
	// see https://martin.baillie.id/wrote/gotchas-in-the-go-network-packages-defaults/#bonus-gomaxprocs-containers-and-the-cfs
	_ "go.uber.org/automaxprocs"
)

type application struct {
	logger *zap.Logger
}

func main() {
	if err := internal.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred: %s\n", err)
		os.Exit(1)
	}
}
