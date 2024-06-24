package http

import (
	"back-platform/app/config"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// newHandler creates a new http.Handler with the routes configured
// @title Back Platform API for Medical Catalog
// @version 1.0
// @description This is the API for the Back Platform Medical Catalog
// @host localhost:3000
// @BasePath /

// @securityDefinitions.apiKey Bearer Token
// @in header
// @name Authorization
func newHandler(_ config.Config) (http.Handler, error) {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

	r.PathPrefix("/docs/v1/swagger/").Handler(httpSwagger.Handler())

	return r, nil
}
