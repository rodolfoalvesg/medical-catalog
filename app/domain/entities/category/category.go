package category

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        int
	PublicID  uuid.UUID
	Name      string
	CreatedBy string
	UpdatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}
