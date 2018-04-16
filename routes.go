package main

import (
	"github.com/labstack/echo"
)

// Routes are all api routes and their handlers
func Routes(server *echo.Echo) {
	server.Static("/", "public")
	server.GET("/api/:yyyy", Year)
	server.GET("/api/:yyyy/:mm", Month)
	server.GET("/api/:yyyy/:mm/:dd", Day)
	server.GET("/api/today", Today)
}
