package specialty

import (
	"back-platform/app/domain/usecases"
	"context"
	"fmt"
)

func (u Usecase) CreateSpecialty(ctx context.Context, input usecases.CreateSpecialtyInput) error {
	const operation = "SpecialtyUsecase.CreateSpecialty"

	if err := u.repository.CreateSpecialty(ctx, input); err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	return nil
}
