package category

import (
	"back-platform/app/gateway/http/category/schema"
	"back-platform/app/gateway/http/rest/responses"
	"fmt"
	"net/http"
)

func (h Handler) GetCategories(r *http.Request) responses.Response {
	const operation = "CategoryHandler.GetSpecialties"

	specialties, err := h.usecase.GetCategories(r.Context())
	if err != nil {
		return responses.InternalServerError(fmt.Errorf("%s: %w", operation, err))
	}

	return responses.Ok(schema.MapToGetCategoriesResponse(specialties))
}
