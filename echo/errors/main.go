package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// IngressError is the return error manipulating ingresses.
type IngressError struct {
	httpCode int
	Code     int    `json:"code"`
	Message  string `json:"message"`
}

func (e IngressError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	fmt.Println("custom error handler is executed.")
	if ie, ok := err.(*IngressError); ok {
		c.JSON(ie.httpCode, echo.Map{
			"success": false,
			"code":    ie.Code,
			"message": ie.Message,
		})
	} else if he, ok := err.(*echo.HTTPError); ok {
		c.JSON(he.Code, echo.Map{
			"success": false,
			"message": he.Message,
		})
	} else {
		c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"message": err.Error(),
		})
	}
}

// Handler is the handler.
func Handler(c echo.Context) error {
	e := &IngressError{
		Code:     401,
		Message:  "test error message",
		httpCode: http.StatusNotFound,
	}
	fmt.Println("Handler executed!")
	return e
}

// Middleware is the middleware.
func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		e := &IngressError{
			Code:     402,
			Message:  "test middleware error message",
			httpCode: http.StatusBadRequest,
		}
		return e
	}
}

func main() {
	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.GET("/error", Handler, Middleware)
	e.Start(":1323")
}
