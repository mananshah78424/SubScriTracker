package db

import (
	"context"
	"subscription_tracker/pkg/application"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UsersModel interface {
	CreateUser(ctx context.Context, app *application.App, users Users) (*Users, error)
	GetAllUsers(ctx context.Context, app *application.App) ([]Users, error)
}

func (u *Users) CreateUser(ctx context.Context, app *application.App, users Users) (*Users, error) {
	// Generate a UUID for the user
	users.ID = uuid.New().String()
	users.CreatedAt = time.Now()
	users.UpdatedAt = time.Now()

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
	ID            string    `bun:"id,pk"`
	Name          string    `bun:"name"`
	Email         string    `bun:"email"`
	Password      string    `bun:"password"`
	CreatedAt     time.Time `bun:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at"`
}
