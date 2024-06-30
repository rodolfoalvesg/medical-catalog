package category

import (
	"back-platform/app/gateway/http/category/schema"
	"back-platform/app/gateway/http/rest/responses"
	"fmt"
	"net/http"
)

// GetCategories returns all categories.
// @Summary Get all categories
// @Description Get all categories
// @Tags Category
// @Security BearerToken
// @Success 200 {object} schema.GetCategoriesResponse "Categories"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Header all {string} RequestID "x-request-id"
// @Router /admin/v1/medical-catalog/categories [get]
func (h Handler) GetCategories(r *http.Request) responses.Response {
	const operation = "CategoryHandler.GetSpecialties"

	specialties, err := h.usecase.GetCategories(r.Context())
	if err != nil {
		return responses.InternalServerError(fmt.Errorf("%s: %w", operation, err))
	}

	return responses.Ok(schema.MapToGetCategoriesResponse(specialties))
}
