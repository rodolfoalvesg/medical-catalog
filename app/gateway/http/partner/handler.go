package partner

import "back-platform/app/domain/usecases"

type Handler struct {
	usecase usecases.PartnerUsecase
}

func NewHandler(usecase usecases.PartnerUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
