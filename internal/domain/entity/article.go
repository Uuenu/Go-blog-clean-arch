package entity

type Article struct {
	ID     string `json:"id,omitempty"`
	Author Author `json:"author,omitempty"`
	Header string `json:"header,omitempty"`
	Text   string `json:"text,omitempty"`
}
