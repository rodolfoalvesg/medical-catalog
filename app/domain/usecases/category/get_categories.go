package category

import (
	"back-platform/app/domain/entities/category"
	"context"
	"fmt"
)

func (u Usecase) GetCategories(ctx context.Context) ([]category.Category, error) {
	const operation = "CategoryUsecase.GetCategories"

	categories, err := u.repository.GetCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	return categories, nil
}
