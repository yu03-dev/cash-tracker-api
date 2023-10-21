package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, e *echo.Echo) {
	publicRouter := e.Group("")
	NewExpenseRouter(db, publicRouter)
	NewCategoryRouter(db, publicRouter)
	NewUserRouter(db, publicRouter)
}
