package login

import (
	"back-platform/app/domain/entities/user"
	"context"

	"github.com/google/uuid"
)

type LoginRepository interface {
	GetPasswd(ctx context.Context, userID uuid.UUID, password string) error
	GetUser(ctx context.Context, email string) (user.User, error)
}
