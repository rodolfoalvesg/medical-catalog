package user

import (
	"back-platform/app/domain/usecases"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, input usecases.CreateUserInput) error
}
