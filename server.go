package main

import (
	"strconv"

	"github.com/google/btree"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Server Instance and data which will be served
type Server struct {
	echo *echo.Echo
	data *btree.BTree
}

// InitServer initializes a new server instance
func InitServer(path string) (Server, error) {
	// New instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Create routes
	Routes(e)

	d, err := loadData(path)
	if err != nil {
		return Server{}, err
	}

	return Server{
		echo: e,
		data: d,
	}, nil
}

// Update reloads the server data
func (s Server) Update(path string) error {

	d, err := loadData(path)
	if err != nil {
		return err
	}
	s.data = d
	return nil
}

// Start starts the listener on the specified port
func (s Server) Start(port int) {
	p := strconv.Itoa(port)
	// Start server
	s.echo.Logger.Fatal(s.echo.Start(":" + p))
}

func loadData(path string) (*btree.BTree, error) {

	verses, err := ImportVerses(path)
	if err != nil {
		return nil, err
	}

	d, err := InitTree(verses)
	if err != nil {
		return nil, err
	}
	return d, nil
}
