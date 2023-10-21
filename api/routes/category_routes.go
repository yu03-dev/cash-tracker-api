package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yu03-dev/cash-tracker-api/api/controller"
	"github.com/yu03-dev/cash-tracker-api/repository"
	"github.com/yu03-dev/cash-tracker-api/usecase"
	"gorm.io/gorm"
)

func NewCategoryRouter(db *gorm.DB, group *echo.Group) {
	cr := repository.NewCategoryRepository(db)
	cc := &controller.CategoryController{
		CategoryUsecase: usecase.NewCategoryUsecase(cr),
	}
	group.GET("/categories", cc.Fetch)
	group.POST("/category", cc.Create)
}
