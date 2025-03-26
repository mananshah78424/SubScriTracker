package server

import (
	"fmt"
	"net/http"
	"subscription_tracker/pkg/application"
	"subscription_tracker/pkg/netflix"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func CreateEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return e
}

func New(app *application.App) (*http.Server, error) {
	e := CreateEcho()
	err := RegisterRoutes(app, e)
	if err != nil {
		fmt.Println("Error registering routes")
		return nil, err
	}
	addr := fmt.Sprintf(":%d", app.Config.ServerPort)
	addr = "localhost" + addr

	s := &http.Server{
		Addr:         addr,
		Handler:      e,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	fmt.Println("Server created on port", app.Config.ServerPort)
	return s, nil

}

func RegisterRoutes(app *application.App, e *echo.Echo) error {
	netflix.RegisterRoutes(e, app)
	return nil
}
