package specialty

import (
	"context"
)

type SpecialtyRepository interface {
	CreateSpecialty(ctx context.Context, specialty Specialty) error
	GetSpecialties(ctx context.Context) ([]Specialty, error)
}
