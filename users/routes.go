package users

import (
	"subscription_tracker/pkg/application"

	"github.com/labstack/echo"
)

type Handler struct {
	app *application.App
}

func RegisterRoutes(e *echo.Echo, app *application.App) {
	// Create a new handler
	h := &Handler{
		app: app,
	}
	group := e.Group("/api/v1")
	group.GET("/users", h.GetUsers)
	group.POST("/user", h.CreateUser)
}
