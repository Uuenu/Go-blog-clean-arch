package entity

type Article struct {
	ID       string `json:"id" bson:"_i, omitempty"`
	AuthorID string `json:"author,omitempty" bson:"author_id,omitempty"`
	Header   string `json:"header,omitempty" bson:"header,omitempty"`
	Text     string `json:"text,omitempty" bson:"text,omitempty"`
}
