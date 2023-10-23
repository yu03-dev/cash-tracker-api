package repository

import (
	"github.com/yu03-dev/cash-tracker-api/domain"
	"gorm.io/gorm"
)

type expenseRepository struct {
	db *gorm.DB
}

func NewGormExpenseRepository(db *gorm.DB) domain.ExpenseRepository {
	return &expenseRepository{db}
}

func (r expenseRepository) Fetch() ([]*domain.Expense, error) {
	expenses := []*domain.Expense{}
	if err := r.db.Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}

func (r expenseRepository) Create(expense *domain.Expense) error {
	return r.db.Create(&expense).Error
}

func (r expenseRepository) GetByUserID(uid int) ([]*domain.Expense, error) {
	expenses := []*domain.Expense{}
	if err := r.db.Find(&expenses, domain.Expense{UserID: uid}).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}

func (r expenseRepository) GetTotalExpenseByUserID(uid int) ([]*domain.TotalExpense, error) {
	totalExpense := []*domain.TotalExpense{}
	err := r.db.Model(&domain.Expense{}).Select("category_id, sum(price) as total").Where("user_id = ?", uid).Group("category_id").Find(&totalExpense).Error
	if err != nil {
		return nil, err
	}
	return totalExpense, nil
}
