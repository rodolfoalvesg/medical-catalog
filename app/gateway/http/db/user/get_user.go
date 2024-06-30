package user

import (
	"back-platform/app/domain/entities/user"
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

func (r *Repository) GetUser(ctx context.Context, email string) (user.User, error) {
	const operation = "UserRepository.GetUser"

	var userStored user.User
	if err := r.db.QueryRow(ctx, _getUserQuery, email).Scan(&userStored.PublicID, &userStored.Username); err != nil {
		if err == sql.ErrNoRows {
			return user.User{}, fmt.Errorf("%s: %w", operation, vos.ErrUserNotFound)
		}

		return user.User{}, fmt.Errorf("%s: %w", operation, err)
	}

	return userStored, nil
}
