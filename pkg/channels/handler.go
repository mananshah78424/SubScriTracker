package channels

import (
	"subscription_tracker/pkg/channels/db"

	"github.com/labstack/echo"
)

func (h *Handler) GetMovies(c echo.Context) error {
	return c.JSON(200, "movies")
}
func (h *Handler) CreateChannel(c echo.Context) error {
	channel, err := validateNewChannelData(c)
	if err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}
	newChannel, err := db.SubmitChannel(c.Request().Context(), h.app, channel)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, newChannel)

}
