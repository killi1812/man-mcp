package man

import "context"

// SearchResult holds an individual match from apropos.
type SearchResult struct {
	Name        string `json:"name"`
	Section     string `json:"section"`
	Description string `json:"description"`
}

// Page holds parsed content and sections of a man page.
type Page struct {
	Command    string            `json:"command"`
	Section    int               `json:"section,omitempty"`
	RawContent string            `json:"raw_content"`
	Sections   map[string]string `json:"sections"`
}

// Client defines the interface for querying system man pages.
type Client interface {
	Search(ctx context.Context, query string, section int) ([]SearchResult, error)
	Fetch(ctx context.Context, command string, section int) (string, error)
}
