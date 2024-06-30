package schema

import (
	"back-platform/app/domain/entities/specialty"

	"github.com/google/uuid"
)

type CreateSpecialtyRequest struct {
	SpecialtyName string `json:"specialtyName"`
}

type GetSpecialtiesResponse struct {
	SpecialtyID   uuid.UUID `json:"specialty_id"`
	SpecialtyName string    `json:"specialty_name"`
	CreatedBy     string    `json:"created_by"`
}

func MapToGetSpecialtiesResponse(specialties []specialty.Specialty) []GetSpecialtiesResponse {
	output := make([]GetSpecialtiesResponse, 0, len(specialties))

	for _, specialty := range specialties {
		s := GetSpecialtiesResponse{
			SpecialtyID:   specialty.PublicID,
			SpecialtyName: specialty.Name,
			CreatedBy:     specialty.CreatedBy,
		}

		output = append(output, s)
	}

	return output
}
