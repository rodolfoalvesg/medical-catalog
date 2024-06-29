package category

import (
	"back-platform/app/domain/usecases"
	"context"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, input usecases.CreateCategoryInput) error
}
