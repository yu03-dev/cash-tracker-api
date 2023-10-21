package usecase

import "github.com/yu03-dev/cash-tracker-api/domain"

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{repo}
}

func (u userUsecase) Fetch() ([]*domain.User, error) {
	return u.repo.Fetch()
}

func (u userUsecase) Create(user *domain.User) error {
	return u.repo.Create(user)
}
