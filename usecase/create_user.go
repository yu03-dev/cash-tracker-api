package usecase

import (
	"github.com/yu03-dev/cash-tracker-api/domain"
	"github.com/yu03-dev/cash-tracker-api/repository"
)

type createUserInteractor struct {
	repo repository.IUserRepository
}

func NewCreateUserInteractor(repo repository.IUserRepository) ICreateUserInteractor {
	return &createUserInteractor{repo: repo}
}

func (i *createUserInteractor) CreateUser(user *domain.User) error {
	return nil
}
