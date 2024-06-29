package http

import (
	"back-platform/app/config"
	categoryUsecase "back-platform/app/domain/usecases/category"
	loginUsecase "back-platform/app/domain/usecases/login"
	specialtyUsecase "back-platform/app/domain/usecases/specialty"
	userUsecase "back-platform/app/domain/usecases/user"
	categoryHandler "back-platform/app/gateway/http/category"
	categoryRepo "back-platform/app/gateway/http/db/category"
	specialtyRepository "back-platform/app/gateway/http/db/specialty"
	userRepository "back-platform/app/gateway/http/db/user"
	loginHandler "back-platform/app/gateway/http/login"
	"back-platform/app/gateway/http/rest"
	specialtyHandler "back-platform/app/gateway/http/specialty"
	userHandler "back-platform/app/gateway/http/user"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger"
)

// newHandler creates a new HTTP handler.
// @title Medical Catalog API
// @version 1.0
// @BasePath /
//
// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization
//
// @description This is a medical catalog API.
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

	userRepo := userRepository.NewRepository(db)
	userUsecase := userUsecase.NewUsecase(userRepo)
	userHandler := userHandler.NewHandler(userUsecase)

	loginUsecase := loginUsecase.NewUsecase(userRepo)
	loginHandler := loginHandler.NewHandler(loginUsecase)

	categoryRepo := categoryRepo.NewRepository(db)
	categoryUsecase := categoryUsecase.NewUsecase(categoryRepo)
	categoryHandler := categoryHandler.NewHandler(categoryUsecase)

	r.Route("/admin/v1/medical-catalog", func(r chi.Router) {
		r.Route("/specialties", func(r chi.Router) {
			r.Post("/", rest.Handle(specialtyHandler.CreateSpecialty))
		})

		r.Route("/categories", func(r chi.Router) {
			r.Post("/", rest.Handle(categoryHandler.CreateCategory))
		})

		r.Route("/login", func(r chi.Router) {
			r.Post("/", rest.Handle(loginHandler.SignIn))
		})
	})

	r.Route("/service/v1/medical-catalog", func(r chi.Router) {
		r.Route("/admin/user", func(r chi.Router) {
			r.Post("/", rest.Handle(userHandler.CreateUser))
		})
	})

	return r, nil
}
