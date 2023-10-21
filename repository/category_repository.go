package repository

import (
	"github.com/yu03-dev/cash-tracker-api/domain"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return &categoryRepository{db}
}

func (r categoryRepository) Fetch() ([]*domain.Category, error) {
	categories := []*domain.Category{}
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r categoryRepository) Create(category *domain.Category) error {
	return r.db.Create(&category).Error
}
