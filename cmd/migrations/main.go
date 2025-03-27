package main

import (
	"fmt"
	"os"
	"subscription_tracker/pkg/utils"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func main() {
	direction := os.Args[1]
	switch direction {
	case "up":
		// Run the up migration
		fmt.Println("Preparing to execute new migrations.")
	case "down":
		// Run the down migration
		fmt.Println("Preparing to rollback migrations.")
	default:
		fmt.Println("Invalid migration direction. Please use 'up' or 'down'.")
	}

	db, err := utils.NewDatabase()
	if err != nil {
		fmt.Println("Error connecting to database")
		return
	}

	defer db.Close()

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		fmt.Println("Error creating driver")
		return
	}

	m, err := migrate.NewWithDatabaseInstance("file://cmd/migrations", "postgres", driver)
	if err != nil {
		fmt.Println("Error creating migration instance")
		return
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			fmt.Println("Error running up migration")
			return
		}
		fmt.Println("Migrations executed successfully")

	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			fmt.Println("Error running down migration")
			return
		}
		fmt.Println("Migrations rolled back successfully")

	}

	fmt.Println("Migration completed successfully")

}
