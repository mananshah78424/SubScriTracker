package application

import (
	"context"
	"subscription_tracker/pkg/utils"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"
)

type App struct {
	Database *bun.DB
}

// New initializes a new application
func New(ctx context.Context, db *bun.DB) (*App, error) {
	db, err := utils.NewDatabase(db)
	if err != nil {
		errMsg := "could not set up the database"
		return nil, errors.Wrap(err, errMsg)
	}

	return &App{
		Database: db,
	}, nil
}
