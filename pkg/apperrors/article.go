package apperrors

import "errors"

var (
	ErrArticleAlreadyExist             = errors.New("article with given email or username already exist")
	ErrArticleNotFound                 = errors.New("article not found")
	ErrArticleNotArchived              = errors.New("article cannot be archived")
	ErrArticleIncorrectEmailOrPassword = errors.New("incorrect email or password")
	ErrArticlePasswordNotGenerated     = errors.New("password generation error")
	ErrArticleIncorrectPassword        = errors.New("incorrect password")
	ErrArticleContextNotFound          = errors.New("article not found in context")
)
