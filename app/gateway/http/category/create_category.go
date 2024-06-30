package category

import (
	"back-platform/app/domain/usecases"
	"back-platform/app/gateway/http/category/schema"
	"back-platform/app/gateway/http/rest/requests"
	"back-platform/app/gateway/http/rest/responses"
	"fmt"
	"net/http"
)

// CreateCategory creates a new category.
// @Summary Create a new category
// @Description Create a new category
// @Description Otherwise it will return HTTP status code 409.
// @Tags Category
// @Security BearerToken
// @Param Body body schema.CreateCategoryRequest true "Category input"
// @Success 200 {object} string "Category created successfully"
// @Failure 400 {object} string "Invalid request body"
// @Failure 409 {object} string "Category already exists"
// @Failure 500 {object} string "Internal server error"
// @Header all {string} RequestID "x-request-id"
// @Router /admin/v1/medical-catalog/categories [post]
func (h Handler) CreateCategory(r *http.Request) responses.Response {
	const operation = "CategoryHandler.CreateCategory"

	var req schema.CreateCategoryRequest
	if err := requests.DecodeBodyJSON(r, &req); err != nil {
		return responses.BadRequest(fmt.Errorf("%s: %w", operation, err), "invalid request body")
	}

	if req.Name == "" {
		return responses.BadRequest(nil, "missing required fields")
	}

	input := usecases.CreateCategoryInput{
		CategoryName: req.Name,
		CreatedBy:    "7e8d250e-3301-11ef-ac6b-a370833da639",
	}

	if err := h.usecase.CreateCategory(r.Context(), input); err != nil {
		return responses.InternalServerError(fmt.Errorf("%s: %w", operation, err))
	}

	return responses.Created()
}
