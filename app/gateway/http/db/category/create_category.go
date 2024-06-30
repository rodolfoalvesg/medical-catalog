package category

import (
	"back-platform/app/domain/entities/category"
	"context"
	"fmt"
)

const _createCategoryQuery = `
	insert into categories (name, created_by, updated_by)
	values ($1, $2, $3)
`

func (r Repository) CreateCategory(ctx context.Context, input category.Category) error {
	const operation = "CategoryRepository.CreateCategory"

	_, err := r.db.Exec(ctx, _createCategoryQuery, input.Name, input.CreatedBy, input.CreatedBy)
	if err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	return nil
}
