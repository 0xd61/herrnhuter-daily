package main

import (
	"io/ioutil"
	"time"
)

type Verse struct {
	Date        time.Time    `xml:Datum`
	Day         time.Weekday `xml:Wtag`
	Sunday      string       `xml:Sonntag`
	Text        string       `xml:Losungstext`
	Reference   string       `xml:Losungsvers`
	Teaching    string       `xml:Lehrtext`
	TeachingRef string       `xml:Lehrtextvers`
	License     string       `xml:License`
}

func fileReader(path string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return bytes
}

func ImportVerses(dirPath string) ([]Verse, error) {
	panic("not implemented")
}
