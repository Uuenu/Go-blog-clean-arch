package apperrors

import "errors"

var (
	ErrSessionNotFound      = errors.New("session not found")
	ErrSessionNotTerminated = errors.New("current session cannot be terminated, use logout instead")
	ErrSessionNotCreated    = errors.New("error occured during session creation")
)
