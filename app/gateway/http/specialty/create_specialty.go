package specialty

import (
	"back-platform/app/domain/entities/specialty"
	"back-platform/app/domain/usecases"
	"back-platform/app/gateway/http/rest/responses"
	"back-platform/app/gateway/http/specialty/schema"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// CreateSpecialty creates a new specialty
// @Summary Create a new specialty
// @Description Create a new specialty
// @Tags Specialty
// @Accept json
// @Produce json
// @Param input body CreateSpecialtyInput true "Specialty input"
// @Success 201 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/medical-catalog/specialties [post]

// Header all {string} RequestID x-request-id
func (h Handler) CreateSpecialty(r *http.Request) responses.Response {
	const operation = "SpecialtyHandler.CreateSpecialty"

	var req schema.CreateSpecialtyInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return responses.BadRequest(fmt.Errorf("%s: %w", operation, err), "invalid request body")
	}

	subjectID := "7e8d250e-3301-11ef-ac6b-a370833da639"

	input := usecases.CreateSpecialtyInput{
		SpecialtyName: req.SpecialtyName,
		CreatedBy:     subjectID,
	}

	if err := h.usecase.CreateSpecialty(r.Context(), input); err != nil {
		if errors.Is(err, specialty.ErrSpecialtyAlreadyExists) {
			return responses.Conflict(fmt.Errorf("%s: %w", operation, err), "specialty already exists")
		}

		return responses.InternalServerError(fmt.Errorf("%s: %w", operation, err))
	}

	return responses.Created()
}
