package specialty

import (
	"back-platform/app/domain/entities/specialty"
	"back-platform/app/domain/usecases"
	"back-platform/app/gateway/http/rest/requests"
	"back-platform/app/gateway/http/rest/responses"
	"back-platform/app/gateway/http/specialty/schema"
	"errors"
	"fmt"
	"net/http"
)

// CreateSpecialty creates a new specialty.
// @Summary Create a new specialty
// @Description Create a new specialty
// @Description Otherwise it will return HTTP status code 409.
// @Tags Specialty
// @Security BearerToken
// @Param Body body schema.CreateSpecialtyRequest true "Specialty input"
// @Success 200 {object} string "Specialty created successfully"
// @Failure 400 {object} string "Invalid request body"
// @Failure 409 {object} string "Specialty already exists"
// @Failure 500 {object} string "Internal server error"
// @Header all {string} RequestID "x-request-id"
// @Router /admin/v1/medical-catalog/specialties [post]
func (h Handler) CreateSpecialty(r *http.Request) responses.Response {
	const operation = "SpecialtyHandler.CreateSpecialty"

	var req schema.CreateSpecialtyRequest
	if err := requests.DecodeBodyJSON(r, &req); err != nil {
		return responses.BadRequest(fmt.Errorf("%s: %w", operation, err), "invalid body json")
	}

	if req.SpecialtyName == "" {
		return responses.BadRequest(fmt.Errorf("%s: %w", operation, specialty.ErrSpecialtyIsEmpty), "specialty name is empty")
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
