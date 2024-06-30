package usecases

import (
	"back-platform/app/domain/entities/category"
	"context"
)

type CategoryUsecase interface {
	CreateCategory(ctx context.Context, input CreateCategoryInput) error
	GetCategories(ctx context.Context) ([]category.Category, error)
}

type CreateCategoryInput struct {
	CategoryName string
	CreatedBy    string
}
