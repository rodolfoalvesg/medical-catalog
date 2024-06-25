package user

import (
	"back-platform/app/domain/entities/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

var _ user.UserRepository = (*Repository)(nil)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}
