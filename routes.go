package main

import (
	"github.com/labstack/echo"
)

// Routes are all api routes and their handlers
func Routes(server *echo.Echo) {
	server.Static("/", "public")
	server.GET("/:yyyy", Year)
	server.GET("/:yyyy/:mm", Month)
	server.GET("/:yyyy/:mm/:dd", Day)
	server.GET("/today", Today)
}
