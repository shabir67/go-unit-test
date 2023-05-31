package service

import (
	"errors"
	"go-unit-test/entity"
	"go-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(Id string) (*entity.Category, error) {
	category := service.Repository.FindById(Id)
	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}
