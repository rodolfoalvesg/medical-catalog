package login

import "back-platform/app/domain/usecases"

var _ usecases.LoginUsecase = (*Usecase)(nil)

type Usecase struct {
	repository LoginRepository
}

func NewUsecase(repository LoginRepository) *Usecase {
	return &Usecase{
		repository: repository,
	}
}
