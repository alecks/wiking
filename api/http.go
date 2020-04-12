package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var e *echo.Echo

var routes []func()

func serverListen() {
	e = echo.New()
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	setRoutes()

	e.Logger.Fatal(e.Start(":80"))
}

func setRoutes() {
	for _, v := range routes {
		v()
	}
}
