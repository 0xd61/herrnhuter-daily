package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Kaitsh/herrnhuter-daily/verses"

	"github.com/labstack/echo"
)

// Handler

// Hello is a dummy handler
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Today returns the verse of today as json
func Today(c echo.Context) error {
	year, month, day := time.Now().Date()
	verse, err := verses.GetVerse(time.Date(year, month, day, 0, 0, 0, 0, time.UTC))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{Error: err.Error()})
	}
	return c.JSON(http.StatusOK, verse)
}

func Day(c echo.Context) error {
	var year, month, day int
	var err error
	if year, err = strconv.Atoi(c.Param("yyyy")); err != nil {
		return c.JSON(http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{Error: "invalid year"})
	}
	if month, err = strconv.Atoi(c.Param("mm")); err != nil {
		return c.JSON(http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{Error: "invalid month"})
	}
	if day, err = strconv.Atoi(c.Param("dd")); err != nil {
		return c.JSON(http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{Error: "invalid day"})
	}

	verse, err := verses.GetVerse(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{Error: err.Error()})
	}
	return c.JSON(http.StatusOK, verse)
}

func Month(c echo.Context) error {
	year, err := strconv.Atoi(c.Param("yyyy"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{Error: "invalid year"})
	}

	month, err := strconv.Atoi(c.Param("mm"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{Error: "invalid month"})
	}
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.UTC)
	data, err := verses.GetRange(startDate, endDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{Error: "internal error"})
	}
	return c.JSON(http.StatusOK, data)
}

func Year(c echo.Context) error {
	year, err := strconv.Atoi(c.Param("yyyy"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{Error: "invalid Date"})
	}
	startDate := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(year+1, time.January, 1, 0, 0, 0, 0, time.UTC)
	data, err := verses.GetRange(startDate, endDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{Error: "internal error"})
	}
	return c.JSON(http.StatusOK, data)

}
