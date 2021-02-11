package cmd

import "time"

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
