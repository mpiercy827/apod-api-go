package health

import (
	"net/http"

	"github.com/labstack/echo"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/health", check)
}

func check(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]bool{"success": true})
}
