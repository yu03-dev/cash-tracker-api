package usecase

import "github.com/yu03-dev/cash-tracker-api/domain"

type categoryUsecase struct {
	repo domain.CategoryRepository
}

func NewCategoryUsecase(repo domain.CategoryRepository) domain.CategoryUsecase {
	return &categoryUsecase{repo}
}

func (u categoryUsecase) Fetch() ([]*domain.Category, error) {
	return u.repo.Fetch()
}

func (u categoryUsecase) Create(category *domain.Category) error {
	return u.repo.Create(category)
}
