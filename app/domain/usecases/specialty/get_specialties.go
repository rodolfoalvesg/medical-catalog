package specialty

import (
	"back-platform/app/domain/entities/specialty"
	"context"
	"fmt"
)

func (u Usecase) GetSpecialties(ctx context.Context) ([]specialty.Specialty, error) {
	const operation = "SpecialtyUsecase.GetSpecialties"

	specialties, err := u.repository.GetSpecialties(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	return specialties, nil
}
