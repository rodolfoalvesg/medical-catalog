package partner

import (
	"back-platform/app/domain/entities/partner"

	"github.com/jackc/pgx/v5/pgxpool"
)

var _ partner.PartnerRepository = (*Repository)(nil)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}
