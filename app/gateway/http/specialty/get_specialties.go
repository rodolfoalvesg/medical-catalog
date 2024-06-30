package specialty

import (
	"back-platform/app/gateway/http/rest/responses"
	"back-platform/app/gateway/http/specialty/schema"
	"fmt"
	"net/http"
)

// GetSpecialties returns all specialties.
// @Summary Get all specialties
// @Description Get all specialties
// @Tags Specialty
// @Security BearerToken
// @Success 200 {object} schema.GetSpecialtiesResponse "Specialties"
// @Failure 500 {object} string "Internal server error"
// @Header all {string} RequestID "x-request-id"
// @Router /admin/v1/medical-catalog/specialties [get]
func (h Handler) GetSpecialties(r *http.Request) responses.Response {
	const operation = "SpecialtyHandler.GetSpecialties"

	specialties, err := h.usecase.GetSpecialties(r.Context())
	if err != nil {
		return responses.InternalServerError(fmt.Errorf("%s: %w", operation, err))
	}

	return responses.Ok(schema.MapToGetSpecialtiesResponse(specialties))
}
