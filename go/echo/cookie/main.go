package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", writeCookieJon)
	e.GET("/cookie/hello", writeCookie, readAllCookies)

	e.Logger.Fatal(e.Start(":1323"))
}

func writeCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "mellon"
	cookie.Path = "/cookie"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}

func readAllCookies(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		for _, cookie := range c.Cookies() {
			fmt.Println(cookie.Name)
			fmt.Println(cookie.Value)
			fmt.Println(cookie.Path)
		}
		return next(c)
	}
}

func writeCookieJon(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "Jon"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}
