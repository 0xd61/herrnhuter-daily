package verses

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

type xmlInput struct {
	MetaData string  `xml:"FreeXml"`
	Data     []Verse `xml:"Losungen"`
}

func fileReader(path string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// ImportVerses loads all verses from a directory and returns them as slice
func ImportVerses(dirPath string) ([]Verse, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	verses := make([]Verse, 0)
	for _, f := range files {
		bytes, err := fileReader(path.Join(dirPath, f.Name()))
		if err != nil {
			return nil, err
		}

		if http.DetectContentType(bytes) != "text/plain; charset=utf-8" {
			fmt.Printf("%v is not a valid xml document\n", f.Name())
			continue
		}

		data := xmlInput{}
		err = xml.Unmarshal(bytes, &data)
		if nil != err {
			return nil, err
		}
		for index := range data.Data {
			data.Data[index].TermsOfUse = "https://www.losungen.de/download/nutzungsbedingungen/"
			data.Data[index].ReferenceSource = "https://www.losungen.de/fileadmin/media-losungen/download/Losungen_Quellenverzeichnis_2018.pdf"
		}
		verses = append(verses, data.Data...)
	}
	return verses, nil
}
