package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	server *echo.Echo
}

func New() *App {
	server := echo.New()
	server.Use(middleware.Recover())
	return &App{
		server: server,
	}
}

func (a App) ConfigureRoutes() {
	a.server.GET("/v1/public/healthy", a.HealthCheck)
}

func (a App) Start() {
	a.ConfigureRoutes()
	a.server.Start(":5000")
}
