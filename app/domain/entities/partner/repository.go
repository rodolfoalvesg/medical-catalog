package partner

import "context"

type PartnerRepository interface {
	CreatePartner(ctx context.Context, partner Partner) error
}
