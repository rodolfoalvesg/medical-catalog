package category

import "back-platform/app/domain/usecases"

type Handler struct {
	usecase usecases.CategoryUsecase
}

func NewHandler(usecase usecases.CategoryUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
