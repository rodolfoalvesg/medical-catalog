package specialty

import (
	"back-platform/app/domain/entities/specialty"

	"github.com/jackc/pgx/v5/pgxpool"
)

var _ specialty.SpecialtyRepository = (*Repository)(nil)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}
