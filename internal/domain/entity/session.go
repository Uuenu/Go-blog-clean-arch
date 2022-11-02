package entity

import "time"

type Session struct {
	ID       string
	AuthorID string
	CreateAt time.Time
}

func NewSession(aid string) (Session, error) {
	//id, err := generate unique string
	// err

	now := time.Now()

	return Session{
		ID:       "",
		AuthorID: aid,
		CreateAt: now,
	}, nil
}
