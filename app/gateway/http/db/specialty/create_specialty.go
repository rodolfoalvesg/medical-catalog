package specialty

import (
	"back-platform/app/domain/entities/specialty"
	"context"
	"fmt"
)

const _createSpecialtyQuery = `
	insert into specialties (name, created_by, updated_by)
	values ($1, $2, $3)
`

func (r Repository) CreateSpecialty(ctx context.Context, specialty specialty.Specialty) error {
	const operation = "SpecialtyRepository.CreateSpecialty"

	_, err := r.db.Exec(ctx, _createSpecialtyQuery, specialty.Name, specialty.CreatedBy, specialty.CreatedBy)
	if err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	return nil
}
