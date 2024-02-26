package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/yu03-dev/cash-tracker-api/usecase"
)

type UserController struct {
	GetUserDetailsInteractor usecase.IGetUserDetailsInteractor
	CreateUserInteractor     usecase.ICreateUserInteractor
}

func (con *UserController) HandleCreateUser(c echo.Context) error {
	return nil
}

func (con *UserController) HandleGetUserDetatils(c echo.Context) error {
	return nil
}
