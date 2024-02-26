package usecase

import (
	"github.com/yu03-dev/cash-tracker-api/domain"
	"github.com/yu03-dev/cash-tracker-api/repository"
)

type getUserDetailsInteractor struct {
	repo repository.IUserRepository
}

func NewGetUserDetailsInteractor(repo repository.IUserRepository) IGetUserDetailsInteractor {
	return &getUserDetailsInteractor{repo: repo}
}

func (i *getUserDetailsInteractor) GetUserDetails(userID int) (*domain.User, error) {
	return nil, nil
}
