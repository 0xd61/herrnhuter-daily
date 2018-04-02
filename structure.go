package main

import (
	"encoding/xml"
	"time"

	"github.com/google/btree"
)

// Verse is the representation of a verse
type Verse struct {
	Date        customTime   `xml:"Datum" json:"date"`
	Day         time.Weekday `json:"weekday"`
	Sunday      string       `xml:"Sonntag" json:"sunday_name"`
	Text        string       `xml:"Losungstext" json:"verse"`
	Reference   string       `xml:"Losungsvers" json:"reference"`
	Teaching    string       `xml:"Lehrtext" json:"teaching"`
	TeachingRef string       `xml:"Lehrtextvers" json:"teaching_reference"`
	License     string       `xml:"License" json:"license"`
}

// Datatype to be able to unmarshall the time format
type customTime struct {
	time.Time
}

func (c *customTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const format = "2006-01-02T15:04:05" // yyyymmdd date format
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(format, v)
	if err != nil {
		return err
	}
	*c = customTime{parse}
	return nil
}

// Less is the compare function for the tree
func (v Verse) Less(than btree.Item) bool {
	return v.Date.Before(than.(Verse).Date.Time)
}

// InitTree initializes a new binary tree with verses
func InitTree(verses []Verse) (*btree.BTree, error) {
	tree := btree.New(2)

	for _, vers := range verses {
		tree.ReplaceOrInsert(vers)
	}
	return tree, nil
}
