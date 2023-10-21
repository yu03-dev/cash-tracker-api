package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string `json:"categoryName" gorm:"not null;default:null"`
	IsExpense    bool   `json:"isExpense" gorm:"not null;default:true"`
}

type CategoryRepository interface {
	Fetch() ([]*Category, error)
	Create(category *Category) error
}

type CategoryUsecase interface {
	Fetch() ([]*Category, error)
	Create(category *Category) error
}
