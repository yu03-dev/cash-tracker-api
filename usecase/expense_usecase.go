package usecase

import "github.com/yu03-dev/cash-tracker-api/domain"

type expenseUsecase struct {
	repo domain.ExpenseRepository
}

func NewExpenseUsecase(repo domain.ExpenseRepository) domain.ExpenseUsecase {
	return &expenseUsecase{repo}
}

func (u expenseUsecase) Fetch() ([]*domain.Expense, error) {
	return u.repo.Fetch()
}

func (u expenseUsecase) Create(expense *domain.Expense) error {
	return u.repo.Create(expense)
}

func (u expenseUsecase) GetByUserID(uid int) ([]*domain.Expense, error) {
	return u.repo.GetByUserID(uid)
}

func (u expenseUsecase) GetTotalExpenseByUserID(uid int) ([]*domain.TotalExpense, error) {
	return u.repo.GetTotalExpenseByUserID(uid)
}
