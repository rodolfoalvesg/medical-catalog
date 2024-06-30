package category

import (
	"back-platform/app/domain/entities/category"
	"context"
	"fmt"
)

const _getCategoriesQuery = `
	SELECT public_id, name, created_by
	FROM categories
`

func (r Repository) GetCategories(ctx context.Context) ([]category.Category, error) {
	const operation = "CategoryRepository.GetCategories"

	rows, err := r.db.Query(ctx, _getCategoriesQuery)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	var categories []category.Category
	for rows.Next() {
		var c category.Category
		if err := rows.Scan(&c.PublicID, &c.Name, &c.CreatedBy); err != nil {
			return nil, err
		}

		categories = append(categories, c)
	}

	return categories, nil
}
