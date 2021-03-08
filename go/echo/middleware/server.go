package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	// Stats is stats :)
	Stats struct {
		Uptime       time.Time      `json:"uptime"`
		RequestCount uint64         `json:"requestCount"`
		Statuses     map[string]int `json:"statuses"`
		mutex        sync.RWMutex
	}
)

// NewStats creates a stats
func NewStats() *Stats {
	return &Stats{
		Uptime:   time.Now(),
		Statuses: map[string]int{},
	}
}

// Process is the middleware function.
func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Process middleware before next()")
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.RequestCount++
		status := strconv.Itoa(c.Response().Status)
		s.Statuses[status]++
		fmt.Println("Process middleware after next()")
		return nil
	}
}

// Handle is the endpoint to get stats.
func (s *Stats) Handle(c echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	fmt.Println("Handle middleware")
	return c.JSON(http.StatusOK, s)
}

// ServerHeader middleware adds a `Server` header to the response.
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		fmt.Println("ServerHeader middleware")
		return next(c)
	}
}

// SetCookie middleware adds a cookie to the response.
func SetCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie := new(http.Cookie)
		cookie.Name = "username"
		cookie.Value = "jon"
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)
		fmt.Println("SetCookie middleware before next()")
		if err := next(c); err != nil {
			c.Error(err)
		}
		fmt.Println("SetCookie middleware after next()")
		return nil
	}

}

func main() {
	e := echo.New()

	// Debug mode
	e.Debug = true

	//-------------------
	// Custom middleware
	//-------------------
	// Stats
	s := NewStats()
	e.Use(s.Process)
	e.GET("/stats", s.Handle, SetCookie) // Endpoint to get stats

	// Server header
	e.Use(ServerHeader)

	// Handler
	e.GET("/", func(c echo.Context) error {
		fmt.Println("/ request")
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Group test
	admin := e.Group("/admin", SetCookie)
	admin.GET("/hello", s.Handle)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
