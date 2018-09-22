package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/mpiercy827/apod-api-go/health"
)

// New creates a new instance of an http.Server
func New() *http.Server {
	e := echo.New()

	health.RegisterRoutes(e)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 4000),
		Handler: e,
	}

	return srv
}
