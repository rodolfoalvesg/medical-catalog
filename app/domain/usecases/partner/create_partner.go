package partner

import (
	"back-platform/app/domain/entities/partner"
	"back-platform/app/domain/usecases"
	"context"
	"fmt"
)

func (u Usecase) CreatePartner(ctx context.Context, input usecases.CreatePartnerInput) error {
	const operation = "PartnerUsecase.CreatePartner"

	partner := partner.Partner{
		Name:         input.PartnerName,
		Phone:        input.Phone,
		CelPhone:     input.CelPhone,
		Email:        input.Email,
		ZipCode:      input.ZipCode,
		Address:      input.Address,
		Number:       input.Number,
		Complement:   input.Complement,
		Reference:    input.Reference,
		Neighborhood: input.Neighborhood,
		City:         input.City,
		UF:           input.UF,
	}

	if err := u.repository.CreatePartner(ctx, partner); err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	return nil
}
