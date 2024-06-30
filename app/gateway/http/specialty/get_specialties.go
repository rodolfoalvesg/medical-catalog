package specialty

import (
	"back-platform/app/gateway/http/rest/responses"
	"fmt"
	"net/http"
)

func (h Handler) GetSpecialties(r *http.Request) responses.Response {
	const operation = "SpecialtyHandler.GetSpecialties"

	specialties, err := h.usecase.GetSpecialties(r.Context())
	if err != nil {
		return responses.InternalServerError(fmt.Errorf("%s: %w", operation, err))
	}

	return responses.Ok(specialties)
}
