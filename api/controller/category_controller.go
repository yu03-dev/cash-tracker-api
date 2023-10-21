package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yu03-dev/cash-tracker-api/domain"
)

type CategoryController struct {
	CategoryUsecase domain.CategoryUsecase
}

func (uc *CategoryController) Fetch(c echo.Context) error {
	categories, err := uc.CategoryUsecase.Fetch()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, categories)
}

func (uc *CategoryController) Create(c echo.Context) error {
	category := domain.Category{}
	if err := c.Bind(&category); err != nil {
		return err
	}
	if err := uc.CategoryUsecase.Create(&category); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, category)
}
