package parser

// Marker represents a status emoji on an entry.
type Marker int

const (
	MarkerAbandoned Marker = iota // :skull:
	MarkerPaid                    // :yen:
	MarkerWIP                     // :construction:
	MarkerStale                   // :ice_cube:
)

// Entry is a single link entry in the README.
type Entry struct {
	Name        string
	URL         string
	Description string
	Raw         string
	Markers     []Marker
	Line        int
}

// Section is a heading with optional entries and child sections.
type Section struct {
	Title    string
	Entries  []Entry
	Children []Section
	Level    int
	Line     int
}

// Document is the parsed representation of the full README.
type Document struct {
	Preamble []string // lines before the first section
	Sections []Section
}
