package partner

import (
	"back-platform/app/domain/entities/partner"
	"back-platform/app/domain/usecases"
)

var _ usecases.PartnerUsecase = (*Usecase)(nil)

type Usecase struct {
	repository partner.PartnerRepository
}

func NewUsecase(repository partner.PartnerRepository) *Usecase {
	return &Usecase{
		repository: repository,
	}
}
