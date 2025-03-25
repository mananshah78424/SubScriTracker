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

func New(app *application.App) *http.Server {
	e := CreateEcho()
	err := RegisterRoutes(app, e)
	if err != nil {
		fmt.Println("Error registering routes")
		return nil
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

	return s

}

func RegisterRoutes(app *application.App, e *echo.Echo) any {
	netflix.RegisterRoutes(e, app)
	return nil
}
