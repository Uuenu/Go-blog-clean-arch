package entity

type Article struct {
	ID       string `json:"id,omitempty" bson:"_id"`
	AuthorID string `json:"author,omitempty" bson:"author_id"`
	Header   string `json:"header,omitempty" bson:"header"`
	Text     string `json:"text,omitempty" bson:"text"`
}
