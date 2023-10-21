package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yu03-dev/cash-tracker-api/domain"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func (uc *UserController) Fetch(c echo.Context) error {
	users, err := uc.UserUsecase.Fetch()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func (uc *UserController) Create(c echo.Context) error {
	user := domain.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := uc.UserUsecase.Create(&user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
