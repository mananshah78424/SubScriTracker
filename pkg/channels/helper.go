package channels

import (
	"fmt"
	"strconv"
	"subscription_tracker/pkg/channels/db"

	"github.com/labstack/echo"
)

func validateNewChannelData(c echo.Context) (db.Channel, error) {
	channelName := c.FormValue("name")
	if channelName == "" {
		fmt.Println("Error: Missing channel name")
		return db.Channel{}, fmt.Errorf("missing channel name")
	}

	channelDescription := c.FormValue("description")
	if channelDescription == "" {
		fmt.Println("Error: Missing channel description")
		return db.Channel{}, fmt.Errorf("missing channel description")
	}

	channelPriceStr := c.FormValue("price")
	if channelPriceStr == "" {
		fmt.Println("Error: Missing channel price")
		return db.Channel{}, fmt.Errorf("missing channel price")
	}

	// Convert price to int
	channelPrice, err := strconv.Atoi(channelPriceStr)
	if err != nil {
		fmt.Println("Error: Invalid channel price")
		return db.Channel{}, fmt.Errorf("invalid channel price")
	}

	channelID := c.FormValue("channel_id")
	if channelID == "" {
		fmt.Println("Error: Missing channel ID")
		return db.Channel{}, fmt.Errorf("missing channel ID")
	}

	// Construct and return the db.Channel object with exported field names
	return db.Channel{
		Name:        channelName,
		Description: channelDescription,
		Price:       channelPrice,
		ChannelID:   channelID,
	}, nil
}
