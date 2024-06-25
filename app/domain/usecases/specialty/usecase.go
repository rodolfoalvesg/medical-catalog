package specialty

import (
	"back-platform/app/domain/entities/specialty"
	"back-platform/app/domain/usecases"
)

var _ usecases.SpecialtyUsecase = (*Usecase)(nil)

type Usecase struct {
	repository specialty.SpecialtyRepository
}

func NewUsecase(repository specialty.SpecialtyRepository) *Usecase {
	return &Usecase{
		repository: repository,
	}
}
