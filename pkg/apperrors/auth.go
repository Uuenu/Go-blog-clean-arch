package apperrors

import "errors"

var (
	ErrAuthAccessDenied = errors.New("access denied")
)
