package login

import (
	"back-platform/app/auth"
	"back-platform/app/domain/usecases"
	"context"
	"fmt"
)

func (u Usecase) SignIn(ctx context.Context, input usecases.SignInInput) (string, error) {
	const operation = "LoginUsecase.SignIn"

	userData, err := u.repository.GetUser(ctx, input.Email)
	if err != nil {
		return "", fmt.Errorf("%s: %w", operation, err)
	}

	if err := u.repository.GetPasswd(ctx, userData.PublicID, input.Password); err != nil {
		return "", fmt.Errorf("%s: %w", operation, err)
	}

	userToken := auth.InputToken{
		PublicID: userData.PublicID,
		Name:     userData.Username,
		Email:    input.Email,
	}

	token, err := auth.GenerateToken(ctx, userToken)
	if err != nil {
		return "", fmt.Errorf("%s: %w", operation, err)
	}

	return token, nil
}
