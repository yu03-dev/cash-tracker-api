package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yu03-dev/cash-tracker-api/api/controller"
	mygorm "github.com/yu03-dev/cash-tracker-api/repository/gorm"
	"github.com/yu03-dev/cash-tracker-api/usecase"
	"gorm.io/gorm"
)

func NewCategoryRouter(db *gorm.DB, group *echo.Group) {
	cr := mygorm.NewGormCategoryRepository(db)
	cc := &controller.CategoryController{
		CategoryUsecase: usecase.NewCategoryUsecase(cr),
	}
	group.GET("/categories", cc.Fetch)
	group.POST("/category", cc.Create)
}
