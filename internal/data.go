package internal

import (
	"encoding/xml"
	"fmt"
	"time"
)

type CalibreBook struct {
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	Author      string   `json:"authors"`
	Sort        string   `json:"author_sort"`
	Languages   []string `json:"languages"`
	Description string   `json:"comments"`
	Cover       string   `json:"cover"`
	Formats     []string `json:"formats"`
	Publisher   string   `json:"publisher"`
	Size        int      `json:"size"`
}

func (c CalibreBook) String() string {
	return fmt.Sprintf("[%04d] %s (%s)", c.Id, c.Title, c.Author)
}

type Opf struct {
	Package  xml.Name `xml:"package"`
	Metadata struct {
		ISBN        string `xml:"identifier"`
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Creator     string `xml:"creator"`
		Date        string `xml:"date"`
		Publisher   string `xml:"publisher"`
		Language    string `xml:"language"`
		Format      string `xml:"format"`
	} `xml:"metadata"`
}

type FSBook struct {
	Id           int
	DirPath      string
	Epub         string
	OriginalEpub string
}

type Index struct {
	Timestamp time.Time `yaml:"timestamp"`
	FSBooks   []FSBook  `yaml:"ebooks"`
}
