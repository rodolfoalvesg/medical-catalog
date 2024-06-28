package usecases

import "context"

type LoginUsecase interface {
	SignIn(ctx context.Context, input SignInInput) (string, error)
}

type SignInInput struct {
	Email    string
	Password string
}

type SignInOutput struct {
	Name     string
	PublicID string
}
