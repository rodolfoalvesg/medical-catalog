package category

import (
	"context"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category Category) error
}
