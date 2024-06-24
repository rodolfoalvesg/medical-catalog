package http

import (
	"back-platform/app/config"
	"net/http"
)

func NewServer(cfg config.Config) (*http.Server, error) {
	handler, err := newHandler(cfg)
	if err != nil {
		return nil, err
	}

	srv := &http.Server{
		Addr:         cfg.HTTP.Address,
		Handler:      handler,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
	}

	return srv, nil
}
