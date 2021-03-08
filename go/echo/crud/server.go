package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

type (
	user struct {
		ID    int    `json:"id"`
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}

	// CustomValidator is the custom validator
	CustomValidator struct {
		validator *validator.Validate
	}
)

// Validate is used to validate struct
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

var (
	users = map[int]*user{}
	seq   = 1
)

//----------
// Handlers
//----------

func responseHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"code":    200,
		"message": "success",
		"data":    c.Get("data"),
	})
}

func createUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := &user{
			ID: seq,
		}
		if err := c.Bind(u); err != nil {
			return err
		}
		if err := c.Validate(u); err != nil {
			return err
		}
		users[u.ID] = u
		seq++
		c.Set("data", u)
		return next(c)
	}
}

func getUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		u := users[id]
		c.Set("data", u)
		return next(c)
	}
}

func updateUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(user)
		if err := c.Bind(u); err != nil {
			return err
		}
		if err := c.Validate(u); err != nil {
			return err
		}
		id, _ := strconv.Atoi(c.Param("id"))
		users[id].Name = u.Name
		users[id].Email = u.Email
		u = users[id]
		c.Set("data", u)
		return next(c)
	}
}

func deleteUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		delete(users, id)
		return next(c)
	}
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/users", responseHandler, createUser)
	e.GET("/users/:id", responseHandler, getUser)
	e.PUT("/users/:id", responseHandler, updateUser)
	e.DELETE("/users/:id", responseHandler, deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
