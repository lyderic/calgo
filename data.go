package main

const (
	cache = "/dev/shm/calgo-cache.json"
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

var valids = []string{
	"12-21",
	"1502",
	"1802",
	"1984",
	"2084",
	"22-11-63",
	"334",
	"658",
	"813",
	"I.G.H.",
	"S.S.S.",
	"meurtre.com",
}

/*
var valids = []string{
	"I",
	"II",
	"III",
	"IV",
	"V",
	"VI",
	"VII",
	"VIII",
	"IX",
	"X",
	"XI",
	"XII",
	"XIII",
	"XIV",
	"XV",
	"XVI",
	"XVIII",
	"XIX",
	"XX",
	"À",
	"J.",
	"M.",
	"Z.",
	"1Q84",
	"SAS",
	"T1",
	"T2",
	"T3",
	"S-300",
	"S.S.S.",
	"12-21",
	"T.A.Z.",
	"Ψ",
	"KGB",
	"CIA",
	"I,",
	"II,",
	"III,",
	"HH7",
	"HH8",
	"HH9",
	"HH10",
	"M.D.O",
	"X.B.T.",
}
*/
