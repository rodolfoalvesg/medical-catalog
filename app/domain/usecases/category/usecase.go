package category

import (
	"back-platform/app/domain/entities/category"
	"back-platform/app/domain/usecases"
)

var _ usecases.CategoryUsecase = (*Usecase)(nil)

type Usecase struct {
	repository category.CategoryRepository
}

func NewUsecase(repository category.CategoryRepository) *Usecase {
	return &Usecase{
		repository: repository,
	}
}
