package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yu03-dev/cash-tracker-api/api/controller"
	"github.com/yu03-dev/cash-tracker-api/repository"
	"github.com/yu03-dev/cash-tracker-api/usecase"
	"gorm.io/gorm"
)

func NewUserRouter(db *gorm.DB, group *echo.Group) {
	ur := repository.NewUserRepository(db)
	uc := &controller.UserController{
		UserUsecase: usecase.NewUserUsecase(ur),
	}
	group.GET("/users", uc.Fetch)
	group.POST("/user", uc.Create)
}
