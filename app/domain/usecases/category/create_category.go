package category

import (
	"back-platform/app/domain/entities/category"
	"back-platform/app/domain/usecases"
	"context"
	"fmt"
)

func (u Usecase) CreateCategory(ctx context.Context, input usecases.CreateCategoryInput) error {
	const operation = "CategoryUsecase.CreateCategory"

	category := category.Category{
		Name:      input.CategoryName,
		CreatedBy: input.CreatedBy,
	}

	if err := u.repository.CreateCategory(ctx, category); err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	return nil
}
