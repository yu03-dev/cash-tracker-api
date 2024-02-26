package domain

import (
	"time"
)

type User struct {
	Name      string
	Email     string
	NickName  string
	Picture   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Category struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FinancialTransaction struct {
	CategoryName string
	Title        string
	FlowType     string
	Amount       int
	Memo         string
	IsAutoSaved  bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type FinancialTransactionGroup struct {
	Name string
	Data []FinancialTransaction
}
