package user

import (
	"back-platform/app/domain/usecases"
	"context"
	"fmt"
)

const _createUserQuery = `
	INSERT INTO users (name, email, passwd)
	VALUES ($1, $2, $3)
	`

func (r Repository) CreateUser(ctx context.Context, input usecases.CreateUserInput) error {
	const operation = "UserRepository.CreateUser"

	_, err := r.db.Exec(context.Background(), _createUserQuery, input.Username, input.Email, input.Password)
	if err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}
	return nil
}
