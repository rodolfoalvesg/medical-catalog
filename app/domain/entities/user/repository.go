package user

import (
	"back-platform/app/domain/usecases"
	"context"
)

type UserRepository struct {
	CreateUser func(ctx context.Context, input usecases.CreateUserInput) error
}
