package structs

import (
	"time"
)

type NewsMessage struct {
	GUID   string
	URL    string
	Sender string
	Title  string
	Text   string
	Created time.Time
	Comments []Comment
}

type Comment struct {
	Author string
	Avatar string
	Text string
	Added time.Time
}