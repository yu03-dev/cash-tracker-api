package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yu03-dev/cash-tracker-api/domain"
)

type ExpenseController struct {
	ExpenseUsecase domain.ExpenseUsecase
}

func (ec *ExpenseController) Create(c echo.Context) error {
	uidStr := c.Param("uid")
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		return err
	}

	expense := domain.Expense{
		UserID: uid,
	}
	if err := c.Bind(&expense); err != nil {
		return err
	}
	if err := ec.ExpenseUsecase.Create(&expense); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, expense)
}

func (ec *ExpenseController) Fetch(c echo.Context) error {
	expenses, err := ec.ExpenseUsecase.Fetch()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, expenses)
}

func (ec *ExpenseController) FetchByUserID(c echo.Context) error {
	uidStr := c.Param("uid")
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		return err
	}
	expenses, err := ec.ExpenseUsecase.GetByUserID(uid)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, expenses)
}

func (ec *ExpenseController) FetchTotalExpenseByUserID(c echo.Context) error {
	uidStr := c.Param("uid")
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		return err
	}
	expneseTotal, err := ec.ExpenseUsecase.GetTotalExpenseByUserID(uid)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, expneseTotal)
}
