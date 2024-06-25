package login

import (
	"back-platform/app/domain/usecases"
	"context"
	"errors"
)

func (u Usecase) SignIn(ctx context.Context, input usecases.SignInInput) (string, error) {
	if input.Username == "" || input.Password == "" {
		return "", errors.New("username or password is empty")
	}

	token, err := u.repository.SignIn(ctx, input)
	if err != nil {
		return "", err
	}

	return token, nil
}
