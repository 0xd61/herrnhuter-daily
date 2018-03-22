package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
