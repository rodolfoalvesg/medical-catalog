package login

import (
	"back-platform/app/domain/usecases"
	"context"
)

type LoginRepository interface {
	GetPasswd(ctx context.Context, userID, password string) error
	GetUser(ctx context.Context, email string) (usecases.SignInOutput, error)
}
