package user

import (
	"back-platform/app/domain/entities/user"
	"back-platform/app/domain/usecases"
	"context"
	"fmt"
)

func (u Usecase) CreateUser(ctx context.Context, input usecases.CreateUserInput) error {
	const operation = "UserUsecase.CreateUser"

	user := user.User{
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
	}

	if err := u.repository.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	return nil
}
