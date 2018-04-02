package main

func main() {

	// Echo instance
	s, err := InitServer("./assets")
	if err != nil {
		return
	}

	s.Start(3333)
}
