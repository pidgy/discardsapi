package api

// Sets represents a collection of cards from a era of a TCG.
type Sets struct {
	Data []SetData `json:"data"`
}

// SetData contains the metadata for individual elements within Sets.
type SetData struct {
	ID    string
	Name  string
	Total int
}
