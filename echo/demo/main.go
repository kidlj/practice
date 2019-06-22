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

	return c.JSON(http.StatusOK, echo.Map{
		"version":  version,
		"hostname": hostname,
	})
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/test", hello)
	e.GET("/hello", hello)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
