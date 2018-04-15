package main

import (
	"github.com/Kaitsh/herrnhuter-daily/verses"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitServer initializes a new server instance
func InitServer(path string) (*echo.Echo, error) {
	// New instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Create routes
	Routes(e)

	err := verses.Update(path)
	if err != nil {
		return nil, err
	}

	return e, nil
}
