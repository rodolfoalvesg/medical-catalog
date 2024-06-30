package usecases

import "context"

type PartnerUsecase interface {
	CreatePartner(ctx context.Context, input CreatePartnerInput) error
}

type CreatePartnerInput struct {
	PartnerName  string
	Phone        string
	CelPhone     string
	Email        string
	ZipCode      string
	Address      string
	Number       string
	Complement   string
	Reference    string
	Neighborhood string
	City         string
	UF           string
}
