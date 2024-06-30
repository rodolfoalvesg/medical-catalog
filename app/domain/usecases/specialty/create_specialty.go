package specialty

import (
	"back-platform/app/domain/entities/specialty"
	"back-platform/app/domain/usecases"
	"context"
	"fmt"
)

func (u Usecase) CreateSpecialty(ctx context.Context, input usecases.CreateSpecialtyInput) error {
	const operation = "SpecialtyUsecase.CreateSpecialty"

	specialty := specialty.Specialty{
		Name:      input.SpecialtyName,
		CreatedBy: input.CreatedBy,
	}

	if err := u.repository.CreateSpecialty(ctx, specialty); err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	return nil
}
