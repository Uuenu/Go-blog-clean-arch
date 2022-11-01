package entity

type Author struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"-"`
	PasswordHash string `json:"-"`
}
