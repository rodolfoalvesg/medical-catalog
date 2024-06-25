package login

import (
	"back-platform/app/domain/usecases"
	"context"
)

type LoginRepository interface {
	SignIn(ctx context.Context, input usecases.SignInInput) (string, error)
}
