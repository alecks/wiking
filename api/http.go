package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var e *echo.Echo

var routes []func()

func serverListen() {
	e = echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} ${method} ${path} - ${latency_human} ${bytes_out}b\n",
	}))
	setRoutes()

	e.Logger.Fatal(e.Start(":80"))
}

func setRoutes() {
	for _, v := range routes {
		v()
	}
}
