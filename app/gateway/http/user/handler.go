package user

import "back-platform/app/domain/usecases"

type Handler struct {
	usecase usecases.UserUsecase
}

func NewHandler(usecase usecases.UserUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
