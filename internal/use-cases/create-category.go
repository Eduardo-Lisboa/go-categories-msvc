package use_cases

import (
	"go-categories-microservice/internal/entities"
	"go-categories-microservice/internal/repositories"
	"log"
)

type createCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewCreateCategoryUseCase(repository repositories.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repository}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)
	if err != nil {
		return err
	}
	//TODO:presit entity to database
	log.Println(category)
	u.repository.Save(category)
	return nil
}
