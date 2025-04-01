package db

import (
	"context"
	"fmt"
	"subscription_tracker/pkg/application"

	"github.com/uptrace/bun"
)

type ChannelModel interface {
	SubmitChannel(ctx context.Context, app *application.App, channel Channel) (*Channel, error)
}

type Channel struct {
	bun.BaseModel `bun:"streaming_catalog,alias:sc"`

	ID          string `bun:"id,pk"`
	Name        string `bun:"name"`
	Description string `bun:"description"`
	Price       int    `bun:"price"`
	ChannelID   string `bun:"channel_id"`
}

func SubmitChannel(ctx context.Context, app *application.App, channel Channel) (*Channel, error) {
	_, err := app.Database.NewInsert().Model(&channel).Exec(ctx)
	if err != nil {
		fmt.Println("Database Insert Error:", err)
		return nil, err
	}
	fmt.Println("Channel inserted successfully:", channel)
	return &channel, nil
}
