package user

import (
	"back-platform/app/domain/entities/user"
	"context"
	"fmt"
)

const _createUserQuery = `
	INSERT INTO users (name, email, passwd)
	VALUES ($1, $2, $3)
	`

func (r Repository) CreateUser(ctx context.Context, user user.User) error {
	const operation = "UserRepository.CreateUser"

	_, err := r.db.Exec(context.Background(), _createUserQuery, user.Username, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}
	return nil
}
