package schema

import (
	"back-platform/app/domain/entities/category"

	"github.com/google/uuid"
)

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type GetCategoriesResponse struct {
	PublicID  uuid.UUID `json:"category_id"`
	Name      string    `json:"category_name"`
	CreatedBy string    `json:"created_by"`
}

func MapToGetCategoriesResponse(categories []category.Category) []GetCategoriesResponse {
	output := make([]GetCategoriesResponse, 0, len(categories))

	for _, category := range categories {
		c := GetCategoriesResponse{
			PublicID:  category.PublicID,
			Name:      category.Name,
			CreatedBy: category.CreatedBy,
		}

		output = append(output, c)
	}

	return output
}
