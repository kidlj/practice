package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Handler
func hello(c echo.Context) error {
	version := os.Getenv("VERSION")
	hostname, err := os.Hostname()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "get hostname error")
	}

	r := c.Request()

	return c.JSON(http.StatusOK, echo.Map{
		"scheme":   c.Scheme(),
		"host":     r.Host,
		"url":      r.URL.Path,
		"realIP":   c.RealIP(),
		"version":  version,
		"hostname": hostname,
	})
}

func notFound(c echo.Context) error {
	return echo.NewHTTPError(http.StatusNotFound, "not found")
}

func serverError(c echo.Context) error {
	return echo.NewHTTPError(http.StatusInternalServerError, "server error")
}

func badRequest(c echo.Context) error {
	return echo.NewHTTPError(http.StatusBadRequest, "bad request")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/400", badRequest)
	e.GET("/404", notFound)
	e.GET("/500", serverError)
	e.GET("/*", hello)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
