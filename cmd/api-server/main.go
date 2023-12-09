package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/edmarfelipe/go-template/internal"
	"github.com/edmarfelipe/go-template/internal/handler"
)

func main() {
	if err := run(); err != nil {
		slog.Error("could not run service", "error", err)
		os.Exit(1)
	}
}

func run() error {
	ct, err := internal.NewContainer()
	if err != nil {
		return fmt.Errorf("failed to create container: %w", err)
	}
	err = handler.NewServer(ct)
	if err != nil {
		return fmt.Errorf("failed to initialize server: %w", err)
	}
	return nil
}
