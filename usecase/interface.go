package usecase

import "github.com/yu03-dev/cash-tracker-api/domain"

type IGetUserDetailsInteractor interface {
	GetUserDetails(userID int) (*domain.User, error)
}

type ICreateUserInteractor interface {
	CreateUser(user *domain.User) error
}
