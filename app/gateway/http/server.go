package http

import (
	"back-platform/app/config"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewServer(cfg config.Config, pgxPool *pgxpool.Pool) (*http.Server, error) {
	handler, err := newHandler(cfg, pgxPool)
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
