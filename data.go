package main

const (
	cache   = "/dev/shm/calgo-cache.json"
	conf    = "/home/lyderic/.calgo.yaml"
	basedir = "/home/lyderic/ebooks/collections/fr"
)

var (
	debug *bool
)

type Book struct {
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

type FSEntry struct {
	Id        int
	Fullpath  string
	Filename  string
	Parentdir string
}

type Opf struct {
	Title       string `xml:"dc:title"`
	Author      string `xml:"dc:creator"`
	Language    string `xml:"dc:language"`
	Publisher   string `xml:"dc:publisher"`
	Description string `xml:"dc:subject"`
}
