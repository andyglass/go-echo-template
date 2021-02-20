package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// RoutesAPIv1 ...
func RoutesAPIv1(e *echo.Group) {
	v1 := e.Group("/v1")
	v1.GET("/info", func(c echo.Context) error {
		return c.String(http.StatusOK, "info")
	})
}
