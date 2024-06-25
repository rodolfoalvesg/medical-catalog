package user

import (
	"back-platform/app/domain/entities/user"
	"back-platform/app/domain/usecases"
)

var _ usecases.UserUsecase = (*Usecase)(nil)

type Usecase struct {
	repository user.UserRepository
}

func NewUsecase(repository user.UserRepository) *Usecase {
	return &Usecase{
		repository: repository,
	}
}
