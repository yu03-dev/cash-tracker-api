package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yu03-dev/cash-tracker-api/api/controller"
	mygorm "github.com/yu03-dev/cash-tracker-api/repository/gorm"
	"github.com/yu03-dev/cash-tracker-api/usecase"
	"gorm.io/gorm"
)

func NewExpenseRouter(db *gorm.DB, group *echo.Group) {
	er := mygorm.NewGormExpenseRepository(db)
	ec := &controller.ExpenseController{
		ExpenseUsecase: usecase.NewExpenseUsecase(er),
	}
	group.GET("/expenses", ec.Fetch)
	group.POST("/expense", ec.Create)
	group.GET("/expenses/:uid", ec.FetchByUserID)
	group.GET("/expenses/total/:uid", ec.FetchTotalExpenseByUserID)
}
