# herrnhuter-daily
[![Go Report Card](https://goreportcard.com/badge/github.com/Kaitsh/herrnhuter-daily)](https://goreportcard.com/report/github.com/Kaitsh/herrnhuter-daily)

This Project serves daily bible verses from the Herrnhuter Brüdergemeinde via a REST API.

## Installation
The software is written in go. To build it a go environment is needed.

```go
    git clone https://github.com/Kaitsh/herrnhuter-daily && cd herrnhuter-daily
    go build -o server
```

To run the server first build the source, then use the command:

```bash
    ./server
```

The server serves the API on `:3333/api`:

Ex.:
```bash
    curl http://localhost:3333/api/today
```

## Routes
The currently available routes are:

```bash
    /
    /api/today
    /yyyy
    /yyyy/mm
    /yyyy/mm/dd
```

## Config
Currently there is no config available. To change settings please change them directly in code. The port is a global variable in the main.go file. Routes can be adjusted in the routes.go file.