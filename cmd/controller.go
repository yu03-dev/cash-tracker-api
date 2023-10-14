package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yu03-dev/cash-tracker-api/pkg/models"
)

func handleGetAllUsers(c echo.Context) error {
	users, err := models.GetAllUsers()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func handleCreateUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := models.CreateUser(&user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)

}

func handleCreateExpense(c echo.Context) error {
	uidStr := c.Param("uid")
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		return err
	}

	expense := models.Expense{
		UserID: uid,
	}
	if err := c.Bind(&expense); err != nil {
		return err
	}
	if err := models.CreateExpense(&expense); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, expense)
}

func handleGetAllExpenses(c echo.Context) error {
	expenses, err := models.GetAllExpenses()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, expenses)
}

func handleGetExpensesByUserID(c echo.Context) error {
	uidStr := c.Param("uid")
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		return err
	}
	expenses, err := models.GetExpensesByUserID(uid)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, expenses)
}

func handleGetTotalExpenseByUserID(c echo.Context) error {
	uidStr := c.Param("uid")
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		return err
	}
	expneseTotal, err := models.GetTotalExpenseByUserID(uid)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, expneseTotal)

}
