package category

import (
	"back-platform/app/domain/usecases"
	"context"
	"fmt"
)

func (u Usecase) CreateCategory(ctx context.Context, input usecases.CreateCategoryInput) error {
	const operation = "CategoryUsecase.CreateCategory"

	if err := u.repository.CreateCategory(ctx, input); err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	return nil
}
