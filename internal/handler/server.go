package handler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/edmarfelipe/go-template/internal"
)

func NewServer(ct internal.Container) error {
	mux := http.NewServeMux()
	mux.Handle("/", &HelloHandler{ct})

	srv := http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}
	slog.Info("Server listening on " + srv.Addr)
	return srv.ListenAndServe()
}
