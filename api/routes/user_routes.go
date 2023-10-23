package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yu03-dev/cash-tracker-api/api/controller"
	mygorm "github.com/yu03-dev/cash-tracker-api/repository/gorm"
	"github.com/yu03-dev/cash-tracker-api/usecase"
	"gorm.io/gorm"
)

func NewUserRouter(db *gorm.DB, group *echo.Group) {
	ur := mygorm.NewGormUserRepository(db)
	uc := &controller.UserController{
		UserUsecase: usecase.NewUserUsecase(ur),
	}
	group.GET("/users", uc.Fetch)
	group.POST("/user", uc.Create)
}
