package main

import (
	"github.com/labstack/echo"
)

// Routes are all api routes and their handlers
func Routes(server *echo.Echo) {
	server.GET("/", hello)
	server.GET("/vers/:yyyy", hello)
	server.GET("/vers/:yyyy/:mm", hello)
	server.GET("/vers/:yyyy/:mm/:dd", hello)
	server.GET("/vers/daily", hello)
	server.GET("/vers/rnd", hello)
	server.GET("/vers", hello)
}
