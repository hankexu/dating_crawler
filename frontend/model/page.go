package model

type SearchResult struct {
	Hits     int64
	Query    string
	Start    int
	PrevFrom int
	NextFrom int
	Items    []interface{}
}
