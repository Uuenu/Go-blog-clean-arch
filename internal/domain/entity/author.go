package entity

import (
	"fmt"
	"go-blog-ca/pkg/utils"
)

const (
	saltSize = 32
)

type Author struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Username     string `json:"username" bson:"username"`
	Email        string `json:"email" bson:"email"`
	Password     string `json:"password" bson:"-"`
	PasswordHash string `json:"-" bson:"password_hash"`
	Salt         []byte `json:"salt" bson:"salt"`
}

// generate password hash
func (a *Author) GeneratePasswordHash() error {

	salt, err := utils.GenerateRandomSalt(saltSize)
	if err != nil {
		return err
	}

	a.Salt = salt

	hashedPassword, err := utils.PasswordHash(a.Password, salt)

	if err != nil {
		return err
	}

	a.PasswordHash = hashedPassword

	return nil
}

// Check if 2 password's match
func (a *Author) DoPasswordMatch(currPassword string) (bool, error) {
	curPswrdHash, err := utils.PasswordHash(currPassword, a.Salt)
	if err != nil {
		return false, fmt.Errorf("failed to compare password. error: %v", err)
	}

	return a.PasswordHash == curPswrdHash, nil
}
