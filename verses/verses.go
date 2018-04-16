package verses

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/google/btree"
)

// Verse is the representation of a verse
type Verse struct {
	Date         customTime `xml:"Datum" json:"date"`
	Day          string     `xml:"Wtag"  json:"weekday"`
	Sunday       string     `xml:"Sonntag" json:"sunday_name"`
	Text         string     `xml:"Losungstext" json:"verse"`
	Reference    string     `xml:"Losungsvers" json:"reference"`
	Teaching     string     `xml:"Lehrtext" json:"teaching"`
	TeachingRef  string     `xml:"Lehrtextvers" json:"teaching_reference"`
	TermsOfUse   string     `json:"terms_of_use"`
	Bibliography string     `json:"bibliography"`
}

// Datatype to be able to unmarshall the time format
type customTime struct {
	time.Time
}

// BTree wrapps the btree package
type BTree struct {
	*btree.BTree
}

// DefaultTree is the default instance of the tree
var DefaultTree = BTree{btree.New(2)}

// Unmarshalling Time is not supported by default
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

// Add adds verses to the binary tree
func (tree *BTree) Add(verses []Verse) error {
	for _, vers := range verses {
		tree.ReplaceOrInsert(vers)
	}
	return nil
}

// GetVerse returns a verse based on a date
func GetVerse(date time.Time) (Verse, error) {
	return DefaultTree.GetVerse(date)
}

// GetVerse returns a verse based on a date
func (tree *BTree) GetVerse(date time.Time) (Verse, error) {
	customDate := customTime{date}
	vers := tree.Get(Verse{Date: customDate})

	if vers == nil {
		return Verse{}, fmt.Errorf("Verse not found")
	}
	return vers.(Verse), nil
}

func GetRange(start, end time.Time) ([]Verse, error) {
	return DefaultTree.GetRange(start, end)
}

func (tree *BTree) GetRange(start, end time.Time) ([]Verse, error) {
	verses := make([]Verse, 0)
	customStart := customTime{start}
	customEnd := customTime{end}

	tree.AscendRange(Verse{Date: customStart}, Verse{Date: customEnd}, func(i btree.Item) bool {
		verses = append(verses, i.(Verse))
		return true
	})

	return verses, nil
}

// Update adds new verses to default tree
func Update(path string) error {
	return DefaultTree.Update(path)
}

// Update adds new verses to tree
func (tree *BTree) Update(path string) error {
	verses, err := ImportVerses(path)
	if err != nil {
		return err
	}

	err = tree.Add(verses)
	if err != nil {
		return err
	}
	return nil
}
