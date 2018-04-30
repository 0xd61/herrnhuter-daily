package main

import (
	"github.com/labstack/echo"
)

// Routes are all api routes and their handlers
func Routes(server *echo.Echo) {
	server.Static("/", "public")
	server.GET("/v/:yyyy", Year)
	server.GET("/v/:yyyy/", Year)
	server.GET("/v/:yyyy/:mm", Month)
	server.GET("/v/:yyyy/:mm/", Month)
	server.GET("/v/:yyyy/:mm/:dd", Day)
	server.GET("/v/:yyyy/:mm/:dd/", Day)
	server.GET("/v", Today)
	server.GET("/v/", Today)
}
