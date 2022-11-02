package entity

type Author struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Username     string `json:"username" bson:"username"`
	Email        string `json:"email" bson:"email"`
	Password     string `json:"password" bson:"-"`
	PasswordHash string `json:"-" bson:"password_hash"`
}
