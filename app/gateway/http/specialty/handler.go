package specialty

import "back-platform/app/domain/usecases"

type Handler struct {
	usecase usecases.SpecialtyUsecase
}

func NewHandler(usecase usecases.SpecialtyUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
