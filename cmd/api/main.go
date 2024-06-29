package main

import (
	"back-platform/app/config"
	httpServer "back-platform/app/gateway/http"
	"context"

	_ "back-platform/docs/swagger"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	pgxPool, err := pgxpool.New(ctx, cfg.DB.ConnectionString())
	if err != nil {
		panic(err)
	}

	defer pgxPool.Close()

	httpSrv, err := httpServer.NewServer(*cfg, pgxPool)
	if err != nil {
		panic(err)
	}

	if err := httpSrv.ListenAndServe(); err != nil {
		panic(err)
	}
}
