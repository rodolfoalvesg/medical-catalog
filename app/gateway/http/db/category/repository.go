package category

import (
	"back-platform/app/domain/entities/category"

	"github.com/jackc/pgx/v5/pgxpool"
)

var _ category.CategoryRepository = (*Repository)(nil)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}
