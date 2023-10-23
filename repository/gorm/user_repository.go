package repository

import (
	"github.com/yu03-dev/cash-tracker-api/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

func (r userRepository) Fetch() ([]*domain.User, error) {
	users := []*domain.User{}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}
