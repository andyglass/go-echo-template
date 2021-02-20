package main

import (
	v1 "go-echo-template/api/v1"
	"net/http"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// GetRouter object
func GetRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())

	prometheus.NewPrometheus("echo", nil).Use(e)

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"ping":"pong"}`)
	})

	api := e.Group("/api")

	v1.RoutesAPIv1(api)

	return e
}
