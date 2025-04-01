package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"subscription_tracker/pkg/utils"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	direction := os.Args[1]

	// Get absolute path to migrations
	migrationPath, err := filepath.Abs("cmd/migrations")
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
	}
	migrationURL := fmt.Sprintf("file://%s", migrationPath)

	db, err := utils.NewDatabase()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error creating postgres driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(migrationURL, "postgres", driver)
	if err != nil {
		log.Fatalf("Error creating migration instance: %v", err)
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Error running up migration: %v", err)
		}
		fmt.Println("Migrations executed successfully")

	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Error running down migration: %v", err)
		}
		fmt.Println("Migrations rolled back successfully")

	default:
		fmt.Println("Invalid migration direction. Please use 'up' or 'down'.")
	}
}
