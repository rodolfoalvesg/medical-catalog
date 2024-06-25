package usecases

import "context"

type LoginUsecase interface {
	SignIn(ctx context.Context, input SignInInput) (string, error)
}

type SignInInput struct {
	Username string
	Password string
}
