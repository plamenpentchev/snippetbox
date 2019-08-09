package models

import (
	"errors"
	"time"
)

//ErrNoRecord ...
var ErrNoRecord = errors.New("models: no matching record found")

//Snippet is a model for an entry in the snippets table in the snippetbox data base
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
