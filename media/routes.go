package media

import (
	"net/http"

	"github.com/labstack/echo"
)

// RegisterRoutes sets up media related routes.
func RegisterRoutes(e *echo.Echo) {
	e.GET("/media", fetch)
}

func fetch(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]bool{"success": true})
}
