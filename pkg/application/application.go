package application

import (
	"context"
	"fmt"
	"subscription_tracker/pkg/config"
	"subscription_tracker/pkg/utils"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"
)

type App struct {
	Database *bun.DB
	Config   *config.Config
}

// New initializes a new application
func New(ctx context.Context) (*App, error) {
	// Set up the database
	db, err := utils.NewDatabase()
	if err != nil {
		errMsg := "could not set up the database"
		return nil, errors.Wrap(err, errMsg)
	}
	fmt.Print("Database connected\n")

	// Application is made up of the database and the configuration
	app := &App{
		Database: db,
		Config:   config.New(),
	}

	return app, nil
}
