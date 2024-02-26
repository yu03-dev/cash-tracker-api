package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/yu03-dev/cash-tracker-api/repository"
	"github.com/yu03-dev/cash-tracker-api/usecase"
	"gorm.io/gorm"
)

func NewUserRouter(db *gorm.DB, group *echo.Group) {
	repo := repository.NewGormUserRepo(db)
	c := UserController{
		GetUserDetailsInteractor: usecase.NewGetUserDetailsInteractor(repo),
		CreateUserInteractor:     usecase.NewCreateUserInteractor(repo),
	}
	group.GET("/user", c.HandleGetUserDetatils)
	group.POST("/user", c.HandleCreateUser)
}

func SetRoutes(db *gorm.DB, e *echo.Echo) {
	publicRouter := e.Group("")
	NewUserRouter(db, publicRouter)
}
