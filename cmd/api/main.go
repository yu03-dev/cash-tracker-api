package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/yu03-dev/cash-tracker-api/database"
	"github.com/yu03-dev/cash-tracker-api/handler"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	handler.SetRoutes(db, e)
	e.Logger.Fatal(e.Start(":3000"))
}
