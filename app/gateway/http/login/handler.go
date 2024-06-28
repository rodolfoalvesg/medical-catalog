package login

import "back-platform/app/domain/usecases"

type Handler struct {
	usecase usecases.LoginUsecase
}

func NewHandler(usecase usecases.LoginUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
