package models

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

func GetAllExpenses() ([]*Expense, error) {
	expenses := []*Expense{}
	if err := db.Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}

func CreateExpense(expense *Expense) error {
	return db.Create(&expense).Error
}

func GetExpensesByUserID(uid int) ([]*Expense, error) {
	expenses := []*Expense{}
	if err := db.Find(&expenses, Expense{UserID: uid}).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}

func GetTotalExpenseByUserID(uid int) ([]*TotalExpense, error) {
	totalExpense := []*TotalExpense{}
	err := db.Model(&Expense{}).Select("category_id, sum(price) as total").Where("user_id = ?", uid).Group("category_id").Find(&totalExpense).Error
	if err != nil {
		return nil, err
	}
	return totalExpense, nil
}
