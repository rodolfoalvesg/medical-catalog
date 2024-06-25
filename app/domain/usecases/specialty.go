package usecases

import "context"

type SpecialtyUsecase interface {
	CreateSpecialty(ctx context.Context, input CreateSpecialtyInput) error
}

type CreateSpecialtyInput struct {
	SpecialtyName string
	CreatedBy     string
}
