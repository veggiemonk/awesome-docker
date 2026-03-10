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
	Markers     []Marker
	Line        int    // 1-based line number in source
	Raw         string // original line text
}

// Section is a heading with optional entries and child sections.
type Section struct {
	Title    string
	Level    int // heading level: 1 = #, 2 = ##, etc.
	Entries  []Entry
	Children []Section
	Line     int
}

// Document is the parsed representation of the full README.
type Document struct {
	Preamble []string // lines before the first section
	Sections []Section
}
