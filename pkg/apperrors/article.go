package apperrors

import "errors"

var (
	ErrArticleAlreadyExist    = errors.New("article with given email or username already exist")
	ErrArticleNotFound        = errors.New("article not found")
	ErrArticleNotArchived     = errors.New("article cannot be archived")
	ErrArticleAccessDenied    = errors.New("access denied")
	ErrArticleContextNotFound = errors.New("article not found in context")
)
