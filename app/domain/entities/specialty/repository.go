package specialty

import (
	"back-platform/app/domain/usecases"
	"context"
)

type SpecialtyRepository interface {
	CreateSpecialty(ctx context.Context, input usecases.CreateSpecialtyInput) error
}
