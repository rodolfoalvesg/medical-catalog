package main

import (
	"back-platform/app/config"
	httpServer "back-platform/app/gateway/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	httpSrv, err := httpServer.NewServer(*cfg)
	if err != nil {
		panic(err)
	}

	if err := httpSrv.ListenAndServe(); err != nil {
		panic(err)
	}
}
