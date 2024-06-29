package usecases

import "context"

type CategoryUsecase interface {
	CreateCategory(ctx context.Context, input CreateCategoryInput) error
}

type CreateCategoryInput struct {
	CategoryName string
	CreatedBy    string
}
