package repository

import "github.com/yu03-dev/cash-tracker-api/domain"

type IUserRepository interface {
	FetchUserByID(userID int) (*domain.User, error)
	CreateUser(user *domain.User) error
}

type ICategoryRepository interface {
	FetchAllCategory(userID int) ([]*domain.Category, error)
	FetchCategoryByName(userID int, name string) (*domain.Category, error)
	CreateCategory(userID int) error
}

type IFinancialTransactionRepository interface {
	FetchAllFinancialTransaction(userID int) ([]*domain.FinancialTransaction, error)
	FetchFinancialTransactionByTitle(userID int, title string) ([]*domain.FinancialTransaction, error)
}

type IFinancialTransactionGroupRepository interface {
	FetchAllFinancialTransactionGroup(userID int) ([]*domain.FinancialTransactionGroup, error)
	FetchFinancialTransactionGroupByName(userID int, name string) ([]*domain.FinancialTransactionGroup, error)
	CreateFinancialTransactionGroup(userID int, group *domain.FinancialTransactionGroup) error
}
