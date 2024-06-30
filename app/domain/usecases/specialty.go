package usecases

import (
	"back-platform/app/domain/entities/specialty"
	"context"
)

type SpecialtyUsecase interface {
	CreateSpecialty(ctx context.Context, input CreateSpecialtyInput) error
	GetSpecialties(ctx context.Context) ([]specialty.Specialty, error)
}

type CreateSpecialtyInput struct {
	SpecialtyName string
	CreatedBy     string
}
