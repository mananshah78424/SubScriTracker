package channels

import (
	"subscription_tracker/pkg/application"

	"github.com/labstack/echo"
)

type Handler struct {
	app *application.App
}

func RegisterRoutes(e *echo.Echo, app *application.App) {
	h := &Handler{
		app: app,
	}
	group := e.Group("/api/v1")
	group.GET("/movies", h.GetMovies)
	group.POST("/channel", h.CreateChannel)
}
