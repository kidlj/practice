package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type IngressError struct {
	Success  bool `json:"success"`
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
		c.JSON(ie.httpCode, ie)
	} else if he, ok := err.(*echo.HTTPError); ok {
		c.String(he.Code, he.Error())
	}
}

func Handle(c echo.Context) error {
	e := &IngressError{
		Success:  false,
		Code:     401,
		Message:  "test error message",
		httpCode: http.StatusNotFound,
	}
	fmt.Println("Handler executed!")
	return e
}

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		e := &IngressError{
			Success:  false,
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

	e.GET("/error", Handle, Middleware)
	e.Start(":1323")
}
