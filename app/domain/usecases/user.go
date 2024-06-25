package usecases

import "context"

type UserUsecase interface {
	CreateUser(ctx context.Context, input CreateUserInput) error
}

type CreateUserInput struct {
	Username string
	Password string
	Email    string
}
