package db

import (
	"context"
	"subscription_tracker/pkg/application"

	"github.com/uptrace/bun"
)

type UsersModel interface {
	CreateUser(ctx context.Context, app *application.App, users Users) (*Users, error)
	GetAllUsers(ctx context.Context, app *application.App) ([]Users, error)
}

func (u *Users) CreateUser(ctx context.Context, app *application.App, users Users) (*Users, error) {
	_, err := app.Database.NewInsert().Model(&users).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *Users) GetAllUsers(ctx context.Context, app *application.App) ([]Users, error) {
	var users []Users
	err := app.Database.NewSelect().Model(&users).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

type Users struct {
	bun.BaseModel `bun:"users,alias:sc"`
	ID            string
	Name          string
	Email         string
	Password      string
}
