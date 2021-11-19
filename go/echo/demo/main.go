package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Handler
func handler(code int) func(echo.Context) error {
	return func(c echo.Context) error {
		version := os.Getenv("VERSION")
		if version == "" {
			version = "unknown"
		}
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}

		r := c.Request()

		return c.JSON(code, echo.Map{
			"scheme":   c.Scheme(),
			"host":     r.Host,
			"url":      r.URL.Path,
			"realIP":   c.RealIP(),
			"version":  version,
			"hostname": hostname,
			"headers":  c.Request().Header,
		})
	}
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/400", handler(http.StatusBadRequest))
	e.GET("/404", handler(http.StatusNotFound))
	e.GET("/500", handler(http.StatusInternalServerError))
	e.GET("/*", handler(http.StatusOK))

	// Start server
	e.Logger.Fatal(e.Start(":9090"))
}
