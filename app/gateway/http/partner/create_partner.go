package partner

import (
	"back-platform/app/domain/entities/partner"
	"back-platform/app/domain/usecases"
	"back-platform/app/gateway/http/partner/schema"
	"back-platform/app/gateway/http/rest/requests"
	"back-platform/app/gateway/http/rest/responses"
	"fmt"
	"net/http"
)

// CreatePartner creates a new partner.
// @Summary Create a new partner
// @Description Create a new partner
// @Tags Partner
// @Security BearerToken
// @Param Body body schema.CreatePartnerRequest true "Partner input"
// @Success 200 {object} string "Partner created successfully"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Header all {string} RequestID "x-request-id"
// @Router /admin/v1/medical-catalog/partners [post]
func (h Handler) CreatePartner(r *http.Request) responses.Response {
	const operation = "PartnerHandler.CreatePartner"

	var req schema.CreatePartnerRequest
	if err := requests.DecodeBodyJSON(r, &req); err != nil {
		return responses.BadRequest(fmt.Errorf("%s: %w", operation, err), "invalid body json")
	}

	if req.PartnerName == "" {
		return responses.BadRequest(fmt.Errorf("%s: %w", operation, partner.ErrPartnerIsEmpty), "partner name is empty")
	}

	input := usecases.CreatePartnerInput{
		PartnerName:  req.PartnerName,
		Phone:        req.Phone,
		CelPhone:     req.CelPhone,
		Email:        req.Email,
		ZipCode:      req.ZipCode,
		Address:      req.Address,
		Number:       req.Number,
		Complement:   req.Complement,
		Reference:    req.Reference,
		Neighborhood: req.Neighborhood,
		City:         req.City,
		UF:           req.UF,
	}

	if err := h.usecase.CreatePartner(r.Context(), input); err != nil {
		return responses.InternalServerError(fmt.Errorf("%s: %w", operation, err))
	}

	return responses.Created()
}
