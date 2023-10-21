package domain

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	UserID     int `json:"userID" gorm:"not null;default:null"`
	Price      int `json:"price" gorm:"not null;default:null"`
	CategoryID int `json:"categoryID" gorm:"not null;default:null"`
}

type TotalExpense struct {
	Total      int `json:"total"`
	CategoryID int `json:"categoryID"`
}

type ExpenseRepository interface {
	Fetch() ([]*Expense, error)
	Create(expense *Expense) error
	GetByUserID(uid int) ([]*Expense, error)
	GetTotalExpenseByUserID(uid int) ([]*TotalExpense, error)
	// GetByID(id int) (*Expense, error)
	// Update(expense *Expense) error
	// Delete(id int) error
}

type ExpenseUsecase interface {
	Fetch() ([]*Expense, error)
	Create(expense *Expense) error
	GetByUserID(uid int) ([]*Expense, error)
	GetTotalExpenseByUserID(uid int) ([]*TotalExpense, error)
	// GetByID(id int) (*Expense, error)
	// Update(expense *Expense) error
	// Delete(id int) error
}
