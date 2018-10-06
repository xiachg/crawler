package model

import (
	"../../engine"
)

type SearchResult struct {
	Hits  int
	Start int
	Items []engine.Item
}
