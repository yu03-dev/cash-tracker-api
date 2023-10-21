package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/yu03-dev/cash-tracker-api/api/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	routes.SetupRoutes(db, e)
	e.Logger.Fatal(e.Start(":3000"))
}
