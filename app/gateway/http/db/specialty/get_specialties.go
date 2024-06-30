package specialty

import (
	"back-platform/app/domain/entities/specialty"
	"context"
	"fmt"
)

const _getSpecialtiesQuery = `
	SELECT
		public_id,
		name,
		created_by
	FROM
		specialties
`

func (r Repository) GetSpecialties(ctx context.Context) ([]specialty.Specialty, error) {
	const operation = "SpecialtyRepository.GetSpecialties"

	rows, err := r.db.Query(ctx, _getSpecialtiesQuery)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	var specialties []specialty.Specialty
	for rows.Next() {
		var s specialty.Specialty
		if err := rows.Scan(&s.PublicID, &s.Name, &s.CreatedBy); err != nil {
			return nil, err
		}

		specialties = append(specialties, s)
	}

	return specialties, nil
}
