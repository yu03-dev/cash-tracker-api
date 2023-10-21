package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"not null;default:null"`
	Email string `json:"email" gorm:"not null;default:null"`
}

type UserRepository interface {
	Fetch() ([]*User, error)
	Create(user *User) error
	// GetByID(id int) (*User, error)
	// Update(user *User) error
	// Delete(id int) error
}

type UserUsecase interface {
	Fetch() ([]*User, error)
	Create(user *User) error
	// GetByID(id int) (*User, error)
	// Update(user *User) error
	// Delete(id int) error
}
