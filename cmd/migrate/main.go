package main

import (
	"flag"
	"fmt"

	"github.com/yu03-dev/cash-tracker-api/database"
	"gorm.io/gorm"
)

func main() {
	migrate := flag.String("migrate", "up", "migration")
	flag.Parse()

	db, err := database.InitDB()

	if err != nil {
		fmt.Println("Database connection error")
	}

	if *migrate == "up" {
		migrateUp(db)
	} else if *migrate == "down" {
		migrateDown(db)
	} else {
		fmt.Println("Invalid command. Use 'up' or 'down'.")
	}

}

func migrateUp(db *gorm.DB) {
	db.AutoMigrate(
		&database.User{},
		&database.Category{},
		&database.FinancialTransaction{},
		&database.FinancialTransactionGroup{},
	)
	fmt.Println("Migration completed successfully.")
}

func migrateDown(db *gorm.DB) {
	db.Migrator().DropTable(
		&database.User{},
		&database.Category{},
		&database.FinancialTransaction{},
		&database.FinancialTransactionGroup{},
	)
	fmt.Println("Tables dropped successfully.")
}
