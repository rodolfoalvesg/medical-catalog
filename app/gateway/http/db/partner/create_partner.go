package partner

import (
	"back-platform/app/domain/entities/partner"
	"context"
	"fmt"
)

const _createPartnerQuery = `
		INSERT INTO partners (
			name,
			phone,
			celphone,
			email,
			zipcode,
			address,
			number,
			complement,
			reference,
			neighborhood,
			city,
			uf,
			created_by,
			updated_by
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
		)
	`

func (r Repository) CreatePartner(ctx context.Context, partner partner.Partner) error {
	const (
		operation = "PartnerRepository.CreatePartner"
		createdBy = "7e8d250e-3301-11ef-ac6b-a370833da639"
	)
	_, err := r.db.Exec(
		ctx,
		_createPartnerQuery,
		partner.Name,
		partner.Phone,
		partner.CelPhone,
		partner.Email,
		partner.ZipCode,
		partner.Address,
		partner.Number,
		partner.Complement,
		partner.Reference,
		partner.Neighborhood,
		partner.City,
		partner.UF,
		createdBy,
		createdBy,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	return nil
}
