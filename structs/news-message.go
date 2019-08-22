package structs

type NewsMessage struct {
	ID     int    `json:"id,omitempty"`
	Sender string `json:"sender"`
	Text   string `json:"text"`
}
