package http

import (
	"back-platform/app/config"
	specialtyUsecase "back-platform/app/domain/usecases/specialty"
	specialtyRepository "back-platform/app/gateway/http/db/specialty"
	"back-platform/app/gateway/http/rest"
	specialtyHandler "back-platform/app/gateway/http/specialty"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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
func newHandler(_ config.Config, db *pgxpool.Pool) (http.Handler, error) {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Route("/docs/v1/medical-catalog", func(r chi.Router) {
		r.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "swagger/index.html", http.StatusMovedPermanently)
		})
		r.Get("/swagger/*", httpSwagger.Handler())
	})

	specialtyRepo := specialtyRepository.NewRepository(db)
	specialtyUsecase := specialtyUsecase.NewUsecase(specialtyRepo)
	specialtyHandler := specialtyHandler.NewHandler(specialtyUsecase)

	r.Route("/admin/v1/medical-catalog", func(r chi.Router) {
		r.Route("/specialties", func(r chi.Router) {
			r.Post("/", rest.Handle(specialtyHandler.CreateSpecialty))
		})
	})

	return r, nil
}
