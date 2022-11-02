package entity

type Article struct {
	ID       string `json:"id,omitempty"`
	AuthorID string `json:"author,omitempty"`
	Header   string `json:"header,omitempty"`
	Text     string `json:"text,omitempty"`
}
