package user

import (
	"back-platform/app/domain/usecases"
	"context"
	"fmt"
)

func (u Usecase) CreateUser(ctx context.Context, input usecases.CreateUserInput) error {
	const operation = "UserUsecase.CreateUser"

	if err := u.repository.CreateUser(ctx, input); err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	return nil
}
