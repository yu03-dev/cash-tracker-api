package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/yu03-dev/cash-tracker-api/pkg/models"
)

func main() {
	err := models.Init(false)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.GET("/users", handleGetAllUsers)
	e.POST("/user", handleCreateUser)
	e.POST("/expense/:uid", handleCreateExpense)
	e.GET("/expenses", handleGetAllExpenses)
	e.GET("/expenses/:uid", handleGetExpensesByUserID)
	e.GET("/expenses/:uid/total", handleGetTotalExpenseByUserID)
	e.Logger.Fatal(e.Start(":3000"))
}
