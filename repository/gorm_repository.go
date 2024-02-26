package repository

import (
	"github.com/yu03-dev/cash-tracker-api/domain"
	"gorm.io/gorm"
)

type gormUserRepo struct {
	db *gorm.DB
}

func NewGormUserRepo(db *gorm.DB) IUserRepository {
	return &gormUserRepo{db: db}
}

func (r *gormUserRepo) FetchUserByID(userID int) (*domain.User, error) {
	return nil, nil
}

func (r *gormUserRepo) CreateUser(user *domain.User) error {
	return nil
}
