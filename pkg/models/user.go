package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"not null;default:null"`
	Email string `json:"email" gorm:"not null;default:null"`
}

func GetAllUsers() ([]*User, error) {
	users := []*User{}
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(user *User) error {
	return db.Create(user).Error
}
