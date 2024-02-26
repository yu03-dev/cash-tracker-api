package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	NickName string
	Picture  string
}

type Category struct {
	gorm.Model
	Name   string
	UserID int
	User   User
}

type FinancialTransaction struct {
	gorm.Model
	CategoryID                  int
	Category                    Category
	Title                       string
	FlowType                    string
	Amount                      int
	Memo                        string
	IsAutoSaved                 bool
	FinancialTransactionGroupID int
	FinancialTransactionGroup   FinancialTransactionGroup
}

type FinancialTransactionGroup struct {
	gorm.Model
	Name string
}

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("TZ"),
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
