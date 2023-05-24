package use_cases

import (
	"go-categories-microservice/internal/entities"
	"go-categories-microservice/internal/repositories"
)

type listCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewListeCategoryUseCase(repository repositories.ICategoryRepository) *listCategoryUseCase {
	return &listCategoryUseCase{repository}
}

func (u *listCategoryUseCase) Execute() ([]*entities.Category, error) {
	categories, err := u.repository.List()
	if err != nil {
		return nil, err
	}
	return categories, nil
}
