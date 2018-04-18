package main

import (
	"strconv"
)

const port = 3333

var version = ""

func main() {

	// Echo instance
	s, err := InitServer("./assets")
	if err != nil {
		return
	}

	s.Logger.Fatal(s.Start(":" + strconv.Itoa(port)))
}
