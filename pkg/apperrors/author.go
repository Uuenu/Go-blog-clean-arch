package apperrors

import "errors"

var (
	ErrAuthorAlreadyExist             = errors.New("author with given email or username already exist")
	ErrAuthorNotFound                 = errors.New("author not found")
	ErrAuthorNotArchived              = errors.New("author cannot be archived")
	ErrAuthorIncorrectEmailOrPassword = errors.New("incorrect email or password")
	ErrAuthorPasswordNotGenerated     = errors.New("password generation error")
	ErrAuthorIncorrectPassword        = errors.New("incorrect password")
	ErrAuthorContextNotFound          = errors.New("author not found in context")
)
