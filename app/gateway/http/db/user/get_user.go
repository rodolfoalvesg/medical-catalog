package user

import (
	"back-platform/app/domain/usecases"
	"back-platform/app/domain/vos"
	"context"
	"database/sql"
	"fmt"
)

const _getUserQuery = `
		SELECT 
			public_id,
			name
		FROM users
		WHERE email = $1
	`

func (r *Repository) GetUser(ctx context.Context, email string) (usecases.SignInOutput, error) {
	const operation = "UserRepository.GetUser"

	var userData usecases.SignInOutput
	if err := r.db.QueryRow(ctx, _getUserQuery, email).Scan(&userData.PublicID, &userData.Name); err != nil {
		if err == sql.ErrNoRows {
			return usecases.SignInOutput{}, fmt.Errorf("%s: %w", operation, vos.ErrUserNotFound)
		}

		return usecases.SignInOutput{}, fmt.Errorf("%s: %w", operation, err)
	}

	return userData, nil
}
