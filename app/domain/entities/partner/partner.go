package partner

import (
	"time"

	"github.com/google/uuid"
)

type Partner struct {
	PublicID     uuid.UUID
	Name         string
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
	CreatedBy    string
	UpdatedBy    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
