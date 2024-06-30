package user

import (
	"back-platform/app/domain/vos"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const _getPasswdQuery = `
		SELECT
			passwd
		FROM users
		WHERE public_id = $1
		`

func (r *Repository) GetPasswd(ctx context.Context, userPublicID uuid.UUID, password string) error {
	const operation = "UserRepository.ValidateUserPassword"

	var storedPassword string
	if err := r.db.QueryRow(ctx, _getPasswdQuery, userPublicID).Scan(&storedPassword); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("%s: %w", operation, err)
		}

		return fmt.Errorf("%s: %w", operation, err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)); err != nil {
		return fmt.Errorf("%s: %w", operation, vos.ErrPasswordMismatch)
	}

	return nil
}
